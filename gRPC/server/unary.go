package main

import (
	"context"

	"github.com/ishansaini194/Projects/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
