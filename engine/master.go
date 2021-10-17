package main

import (
	"context"
	pb "github.com/vynaloze/mapreduce/engine/ping_pong"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPingPongServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PongReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PongReply{Message: "Pong"}, nil
}

func serveMaster(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingPongServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
