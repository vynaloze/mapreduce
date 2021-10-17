package controller

import (
	"context"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type registrarServer struct {
	pb.UnimplementedRegistrarServer
	workers chan<- Worker
}

func (s *registrarServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	w := Worker{Addr: in.GetAddr()}
	log.Printf("Received Register: %v\n", w)
	s.workers <- w
	return &pb.RegisterReply{}, nil
}

func serveRegistrar(addr string) <-chan Worker {
	workers := make(chan Worker, WorkerLimit)
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterRegistrarServer(s, &registrarServer{workers: workers})
		log.Printf("registrarServer listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return workers
}
