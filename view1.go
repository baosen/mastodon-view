package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	pb "github.com/baosen/mastodon_view/mastodon"

	"google.golang.org/grpc"
)

// Views the updates from the puller.
func main() {
	// Connect to the puller.
	connection, err := grpc.Dial(fmt.Sprintf("%s:50051", os.Args[1]), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	// Setup client.
	client := pb.NewExampleServiceClient(connection)

	type PageData struct {
		Title   string
		Content string
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Talk to the puller.
		res, err := client.SendMessage(context.Background(), &pb.MessageRequest{Message: "Hello!"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		// Set the content for your template.
		content := PageData{
			Content: res.GetReply(),
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

	port := ":8081"
	log.Printf("Starting view1 at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
