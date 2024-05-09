package main

import (
	"context"
	"log"
	"net"

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
	pb.RegisterExampleServiceServer(grpcServer, &server{})

	// Start your engines!
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Subscribe subscribes to the puller.
func (s *server) Subscribe(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Reply: req.Message}, nil
}

type server struct {
	pb.ExampleServiceServer
}
