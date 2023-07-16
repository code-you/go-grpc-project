package main

import (
	"context"
	pb "github.com/code-you/go-grpc-project/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error streaming: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending: %v", err)
		}
		log.Printf("Send the request	with the name : %s", name)
		time.Sleep(time.Second * 2)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	log.Printf("%v", res.Messages)
}
