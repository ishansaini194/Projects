package main

import (
	"log"
	"time"

	"github.com/ishansaini194/Projects/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names %v", req.Names)
	for _, names := range req.Names {
		res := &proto.HelloResponse{
			Message: "Hello" + names,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
