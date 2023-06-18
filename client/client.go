package main

import (
	"context"
	"fmt"
	"grpc-test/protos"
	"log"
	"time"
)

func main() {
	// create a connection pool
	pool, err := NewChatServerPool(2, &ServerDetails{Address: ":9000"})
	if err != nil {
		log.Fatalf("error in creating pool to Chat Server. error is: [%s]", err)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Printf("Calling Chat server for: [%d]\n", i)
			callChatServer(pool)
			fmt.Printf("Call completed for Chat server for: [%d]\n", i)
		}(i)
	}

	time.Sleep(10 * time.Minute)
}

func callChatServer(pool *ChatGrpcConnectionPool) {
	conn, err := pool.Get()
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