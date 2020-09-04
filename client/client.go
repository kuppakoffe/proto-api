package main

import (
	"context"
	"fmt"
	greeting_pb "github.com/kuppakoffe/proto-api/api/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err:=grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while connecting to server: %v", err)
	}
	defer conn.Close()

	c:=greeting_pb.NewGreetingClient(conn)
	resp , err:= c.GreetingGet(context.Background(), &greeting_pb.GreetRequest{Name: "Sumit"})
	if err != nil {
		log.Fatalf("Error while getting response: %v", err)
	}
	
	fmt.Println(resp.Message)
}