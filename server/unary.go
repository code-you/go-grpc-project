package main

import (
	"context"
	pb "github.com/code-you/go-grpc-project/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
