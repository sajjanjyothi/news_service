package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/sajjan_jyothi/user_service/user_service_grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var userDatabase = map[string][]string{
	"user1": []string{"sports", "entertainment", "live"},
	"user2": []string{"sports", "entertainment"},
	"user3": []string{"sports", "live"},
	"user4": []string{"sports"},
	"user5": []string{"sports", "entertainment", "live"},
}

func (s *server) GetUserDetails(ctx context.Context, request *pb.UserRequest) (*pb.UserReply, error) {
	log.Println("User details query for ", request.GetUserId())
	tags, err := userDatabase[request.GetUserId()]
	if err == false {
		return nil, errors.New("Cannot find user id")
	}
	return &pb.UserReply{UserId: request.GetUserId(), Tags: tags}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("GRPC server listening on port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
