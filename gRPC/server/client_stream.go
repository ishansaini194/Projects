package main

import (
	"io"
	"log"

	"github.com/ishansaini194/Projects/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
