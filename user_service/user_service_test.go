package main

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/sajjan_jyothi/user_service/user_service_grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func getUserTags(userID string) ([]string, error) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Cannot connect to user_service %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //Adding one second context timeout
	defer cancel()
	r, err := c.GetUserDetails(ctx, &pb.UserRequest{UserId: userID})
	if err != nil {
		log.Printf("could not user details: %v", err)
		return nil, err
	}
	return r.GetTags(), nil
}

func TestUserTags(t *testing.T) {
	tags, _ := getUserTags("user1")

	if len(tags) != 3 {
		t.Error("Error in getting value for ", "user1")
	}

}
