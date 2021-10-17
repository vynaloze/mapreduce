package main

import (
	"context"
	pb "github.com/vynaloze/mapreduce/engine/ping_pong"
	"google.golang.org/grpc"
	"log"
	"time"
)

func serveWorker(masterAddr string) {
	conn, err := grpc.Dial(masterAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingPongClient(conn)

	for {
		request(c)
		time.Sleep(time.Second)
	}
}

func request(c pb.PingPongClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{Message: "Ping"})
	if err != nil {
		log.Fatalf("could ping: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
