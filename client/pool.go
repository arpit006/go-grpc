package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
)

type ChatGrpcConnectionPool struct {
	clients map[int]*grpc.ClientConn
}

type ServerDetails struct {
	 Address string
}

func NewChatServerPool(poolSize int, details *ServerDetails) (*ChatGrpcConnectionPool, error) {
	c := make(map[int]*grpc.ClientConn)
	if poolSize <= 0 {
		poolSize = 1
	}
	for i := 1; i <= poolSize; i++ {
		conn, err := grpc.Dial(details.Address, grpc.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("error in dialing to server: [%+v]. error is: [%s]", details, err)
		}
		log.Printf("Establishing connection to server: [%+v]\n", details)
		c[i] = conn
	}
	return &ChatGrpcConnectionPool{
		clients: c,
	}, nil
}

// Get returns a connection in random from available connection pool
func (pool *ChatGrpcConnectionPool) Get() (*grpc.ClientConn, error) {
	if pool == nil || pool.clients == nil || len(pool.clients) == 0 {
		return nil, fmt.Errorf("no clients available to connect to Pool server, Please initialise the pool first")
	}
	r := rand.Intn(len(pool.clients))
	if r <= 0 || r > len(pool.clients) {
		r = 1
	}
	log.Printf("Returning client: [%d]", r)
	return pool.clients[r], nil
}
