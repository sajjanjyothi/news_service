package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "sajjan.com/sajjan_jyothi/news_article_service/user_service_grpc"
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

func TestRestEndPointValid(t *testing.T) {
	type request struct {
		UserID string
		Tags   []string
	}
	restRequest := request{UserID: "user1", Tags: []string{"sports", "live"}}
	reqBody, err := json.Marshal(restRequest)

	if err != nil {
		t.Error("Request body preparation failed", err)
	}
	resp, _ := http.Post("http://localhost:8000/api/v1/news_list/", "application/json", bytes.NewBuffer(reqBody))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error("news list rest api failed")
	}
	rest_request = Request{UserID: "user10", Tags: []string{"sports", "live"}}
	resp, _ = http.Post("http://localhost:8000/api/v1/news_list/", "application/json", bytes.NewBuffer(reqBody))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Error("news list rest api failed")
	}
}

func TestRestEndPointInValid(t *testing.T) {
	type request struct {
		UserID string
		Tags   []string
	}
	restRequest := request{UserID: "user10", Tags: []string{"sports", "live"}}
	reqBody, err := json.Marshal(restRequest)
	if err != nil {
		t.Error("Request body preparation failed", err)
	}
	resp, _ := http.Post("http://localhost:8000/api/v1/news_list/", "application/json", bytes.NewBuffer(reqBody))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatal("news list rest api failed")
		t.FailNow()
	}
}
