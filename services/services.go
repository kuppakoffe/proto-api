package services

import (
	"fmt"
	"log"
	"net"

	greeting "github.com/kuppakoffe/proto-api/api/proto"
	"google.golang.org/grpc"
)

type server struct{}

func Start() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}

	s := grpc.NewServer()
	greeting.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed tp server: %v", err)
	}
	fmt.Println("started grpc server on port 50051")
}
