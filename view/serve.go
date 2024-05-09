package view

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/baosen/mastodon_view/mastodon"
	pb "github.com/baosen/mastodon_view/mastodon"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

// Views the updates from the puller.
func Serve(title string, port string) {
	// Serve the frontend.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type PageData struct {
			Title string
			Port  string
		}

		// Set the content for your template.
		content := PageData{
			Title: title,
			Port:  port,
		}

		// Parse the HTML template file.
		template, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the provided data and write the output to the response
		err = template.Execute(w, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Connect to the puller.
	connection, err := grpc.Dial(fmt.Sprintf("%s:50051", os.Args[1]), grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	} else {
		defer connection.Close()

		// Setup client.
		client := pb.NewPullerServiceClient(connection)

		// An endpoint for the frontend to subscribe to updates from the puller.
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
			connection, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Println(err)
				return
			}
			defer connection.Close()

			for {
				// Talk to the puller.
				res, err := client.Subscribe(context.Background(), &mastodon.Empty{})
				if err != nil {
					log.Fatalf("did not get an reply: %v", err)
				}

				// Publish the message from the puller to the frontend.
				if err := connection.WriteMessage(websocket.TextMessage, []byte(res.GetReply())); err != nil {
					log.Println(err)
					return
				}

				time.Sleep(time.Duration(1) * time.Second)
			}
		})
	}

	log.Printf("Starting %s at port %s\n", title, port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
