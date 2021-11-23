package controller

import (
	"context"
	"github.com/vynaloze/mapreduce/api"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

func (c *controller) RegisterNewExecutable(e *api.MapReduceExecutable) {
	log.Printf("received new executable to register")
	c.registry.mux.RLock()
	defer c.registry.mux.RUnlock()
	c.registry.executables = append(c.registry.executables, e)
}

type registryServer struct {
	pb.UnimplementedRegistryServer
	workers chan<- Worker

	executables []*api.MapReduceExecutable
	mux         sync.RWMutex
}

func newRegistryServer() *registryServer {
	return &registryServer{executables: make([]*api.MapReduceExecutable, 0)}
}

func (s *registryServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	missingExecutables := s.getMissingExecutables(in)
	if len(missingExecutables) > 0 {
		log.Printf("sending executables %+v to worker %s", missingExecutables, in.GetUuid())
		for _, idx := range missingExecutables {
			// TODO send executalbe
			log.Println(idx)
		}
	}

	w := Worker{Uuid: in.GetUuid(), Addr: in.GetAddr(), ExpiresOn: time.Now().Add(time.Second * time.Duration(in.GetTtlSeconds()))}
	//log.Printf("Received Register: %v\n", w)
	s.workers <- w
	return &pb.RegisterReply{}, nil
}

func (s *registryServer) serve(lis net.Listener, serv *grpc.Server) <-chan Worker {
	workers := make(chan Worker, WorkerLimit)
	s.workers = workers
	go func() {
		defer close(workers)
		pb.RegisterRegistryServer(serv, s)
		log.Printf("registryServer listening at %v", lis.Addr())
		if err := serv.Serve(lis); err != nil {
			log.Fatalf("failed to serve registryServer: %v", err)
		}
	}()
	return workers
}

func (s *registryServer) getMissingExecutables(in *pb.RegisterRequest) []int {
	s.mux.Lock()
	defer s.mux.Unlock()

	missing := make([]int, 0)

Outer:
	for i := range s.executables {
		for _, m := range s.executables[i].GetMappers() {
			if !contains(m, in.GetMappers()) {
				missing = append(missing, i)
				continue Outer
			}
		}
		for _, r := range s.executables[i].GetReducers() {
			if !contains(r, in.GetReducers()) {
				missing = append(missing, i)
				continue Outer
			}
		}
	}
	return missing
}

func contains(s string, slice []string) bool {
	for _, ss := range slice {
		if ss == s {
			return true
		}
	}
	return false
}
