package worker

import (
	"context"
	"fmt"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run(addr, masterAddr string) {
	err := register(addr, masterAddr)
	if err != nil {
		log.Fatal(err)
	}
}

func register(addr, masterAddr string) error {
	conn, err := grpc.Dial(masterAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRegistryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Register(ctx, &pb.RegisterRequest{Addr: addr})
	if err != nil {
		return fmt.Errorf("could not register: %v", err)
	}
	return nil
}
