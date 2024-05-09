package main

import (
	"context"
	"fmt"
	"log"
	"net"

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

	// Start your engines!
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var count = 0

// Subscribe subscribes to the puller.
func (s *server) Subscribe(ctx context.Context, empty *mastodon.Empty) (*pb.Reply, error) {
	count += 1
	return &pb.Reply{Reply: fmt.Sprintf("%d\n", count)}, nil
}

type server struct {
	pb.PullerServiceServer
}
