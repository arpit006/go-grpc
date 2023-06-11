package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test/protos"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect to server on 9000. [%s]", err)
	}

	defer conn.Close()

	c := protos.NewChatServiceClient(conn)
	msg := &protos.Message{
		Body: "Server! Are you there??",
	}

	resp, err := c.SayHello(context.Background(), msg)
	if err != nil {
		log.Fatalf("error received from server. Error is: [%s]\n", err)
	}
	log.Printf("Response: [%s]", resp.Body)
}