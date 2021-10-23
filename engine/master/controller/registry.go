package controller

import (
	"context"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type registryServer struct {
	pb.UnimplementedRegistryServer
	workers chan<- Worker
}

func (s *registryServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	w := Worker{Addr: in.GetAddr()}
	log.Printf("Received Register: %v\n", w)
	s.workers <- w
	return &pb.RegisterReply{}, nil
}

func serveRegistry(lis net.Listener, s *grpc.Server) <-chan Worker {
	workers := make(chan Worker, WorkerLimit)
	go func() {
		pb.RegisterRegistryServer(s, &registryServer{workers: workers})
		log.Printf("registryServer listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve registryServer: %v", err)
		}
	}()
	return workers
}
