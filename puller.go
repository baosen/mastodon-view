package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/baosen/mastodon_view/mastodon"
	pb "github.com/baosen/mastodon_view/mastodon"
	"google.golang.org/grpc"
)

// Read a stream from a Mastodon-server and serve it over gRPC.
func main() {
	log.Printf("Start puller")

	// Setup a gRPC-server that pulls from Mastodon.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPullerServiceServer(grpcServer, &server{})

	updates = make(chan string)
	go func() {
		var count = 0
		for {
			count += 1
			updates <- fmt.Sprintf("%d\n", count)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	// Start your engines!
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Subscribe subscribes to the puller.
func (s *server) Subscribe(empty *mastodon.Empty, stream mastodon.PullerService_SubscribeServer) error {
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
