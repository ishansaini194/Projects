package main

import (
	"context"
	"log"
	"time"

	"github.com/ishansaini194/Projects/proto"
)

func SayHelloClientStreaming(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send name %v", err)
	}

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Send the request with name %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error While receving %v", err)
	}
	log.Printf("%v", res.Messages)
}
