package main

import (
	"log"

	"github.com/ishansaini194/Projects/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Didn't connect %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	names := &proto.NamesList{
		Names: []string{"Ishan", "Tushar", "Uday"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
	SayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
