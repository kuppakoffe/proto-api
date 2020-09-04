package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	greeting "github.com/kuppakoffe/proto-api/api/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/kuppakoffe/proto-api/services"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	flag.Parse()
	go services.Start()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := greeting.RegisterGreetingHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		log.Fatalf("Error while starting rest endpoint: %v", err)
	}
	fmt.Println("Running rest endpoint")

	http.ListenAndServe(":8080", mux)
}
