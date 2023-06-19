package main

import (
	"context"

	pb "github.com/gitaepark/go_grpc_basic/proto"
)



func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}