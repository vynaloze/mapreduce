package master

import (
	"context"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type registrarServer struct {
	pb.UnimplementedRegistrarServer
	workers chan<- string
}

func (s *registrarServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	log.Printf("Received: %v", in.GetAddr())
	s.workers <- in.GetAddr()
	return &pb.RegisterReply{}, nil
}

func serveRegistrar(addr string, workers chan<- string) {
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
}
