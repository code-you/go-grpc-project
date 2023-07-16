package main

import (
	pb "github.com/code-you/go-grpc-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const port = ":8080"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error dialing: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error closing connection: %v", err)
		}
	}(conn)

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Aditya", "Neeraj", "Singh", "Patel", "Ji"},
	}
	// unary call
	//callSayHello(client)
	//server streaming
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
