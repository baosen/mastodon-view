package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/baosen/mastodon_view/mastodon"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

// Read a stream from a Mastodon-server and serve it over gRPC.
func main() {
	log.Printf("Start puller")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up Mastodon client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})

	// Start streaming public timeline
	stream, err := client.StreamingPublic(context.Background(), true)
	if err != nil {
		log.Fatal(err)
	}

	// Setup a gRPC-server that pulls from Mastodon.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPullerServiceServer(grpcServer, &server{})

	updates = make(chan string)
	go func() {
		for {
			select {
			case event := <-stream:
				switch event := event.(type) {
				case *mastodon.UpdateEvent:
					updates <- event.Status.Content
				case *mastodon.ErrorEvent:
					updates <- event.Error()
				}
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	// Start your engines!
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Subscribe subscribes to the puller.
func (s *server) Subscribe(empty *pb.Empty, stream pb.PullerService_SubscribeServer) error {
	select {
	case update := <-updates:
		stream.Send(&pb.Reply{Reply: update})
		return nil
	}
}

type server struct {
	pb.PullerServiceServer
}

var updates chan string
