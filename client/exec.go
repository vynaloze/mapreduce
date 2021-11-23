package client

import (
	"context"
	"flag"
	"fmt"
	"github.com/vynaloze/mapreduce/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func (c *client) registerThisExecutable() error {
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	executableBytes, err := os.ReadFile(executable)

	msg := api.MapReduceExecutable{
		Mapper:     c.mapper.Name,
		Reducer:    c.reducer.Name,
		Executable: executableBytes,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err = c.mapReduceRegistryClient.Register(ctx, &msg)
	return err
}

func startMapReduceServer(c *client) {
	var port int
	flag.IntVar(&port, "mapreduce-port", 0, "port to serve mapreduce on")
	flag.Parse()

	if port != 0 {
		log.Printf("starting mapreduce on %d", port)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		api.RegisterMapReduceServer(s, &mapReduceServer{client: c})
		log.Printf("mapReduceServer listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
