package main

import (
	"context"
	"fmt"
	grpcpool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"grpc-test/protos"
	"log"
	"time"
)

func main() {

	factory := func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Client could not connect to server on 9000. [%s]", err)
		}
		log.Printf("Connected to Server at : [%s]\n", ":9000")
		return conn, err
	}

	pool, err := grpcpool.New(factory, 5, 10, time.Second)
	if err != nil {
		log.Fatalf("Failed to create gRPC pool: %v", err)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Printf("Calling Chat server for: [%d]\n", i)
			callChatServer(pool)
			fmt.Printf("Call completed for Chat server for: [%d]", i)
		}(i)
	}

	time.Sleep(10 * time.Minute)
}

func callChatServer(pool *grpcpool.Pool) {
	ctx := context.Background()
	conn, err := pool.Get(ctx)
	defer conn.Close()
	if err != nil {
		msg := "failed to connect to worker"
		log.Printf("Error: [%s], msg: [%s]\n", err, msg)
		return
	}
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
