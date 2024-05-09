package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/baosen/mastodon_view/mastodon"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the puller.
	connection, err := grpc.Dial(fmt.Sprintf("%s:50051", os.Args[1]), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	// Setup client.
	client := pb.NewExampleServiceClient(connection)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Talk to the puller.
		res, err := client.SendMessage(context.Background(), &pb.MessageRequest{Message: "Hello!"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", res.GetReply())
		http.ServeFile(w, r, "index.html")
	})

	port := ":8081"
	log.Printf("Starting view1 at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
