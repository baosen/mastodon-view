package streamer_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/baosen/mastodon_view/mastodon"
	"github.com/baosen/mastodon_view/streamer"
	"google.golang.org/grpc"
)

func TestMastodonStream(t *testing.T) {
	go func() {
		streamer.Stream("../.env")
	}()

	// Connect to the puller.
	var connection *grpc.ClientConn
	var err error
	for retries := 0; retries < 30; retries++ {
		connection, err = grpc.Dial(fmt.Sprintf("localhost:50051"), grpc.WithInsecure())
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	// Setup client.
	client := pb.NewPullerServiceClient(connection)

	// Get the subscribe client to talk to the puller.
	c, err := client.Subscribe(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("did not get the client: %v", err)
		return
	}

	// Receive reply.
	_, err = c.Recv()
	if err != nil {
		t.Fatalf("did not receive reply: %v", err)
	}
}
