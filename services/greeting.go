package services

import (
	"fmt"
	greeting "github.com/kuppakoffe/proto-api/api/proto"
	"context"
)



func (s *server) GreetingGet(ctx context.Context, req *greeting.GreetRequest) (*greeting.GreetResponse, error) {
	name :=req.GetName()
	message:= fmt.Sprintf("Hello %s", name)
	return &greeting.GreetResponse{
		Message: message}, nil
}