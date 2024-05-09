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
	log.Printf("Start mastodon_puller")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterExampleServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.ExampleServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Reply: req.Message}, nil
}
