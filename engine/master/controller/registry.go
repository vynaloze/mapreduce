package controller

import (
	"context"
	"github.com/vynaloze/mapreduce/api"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func (c *controller) RegisterNewExecutable(e *api.MapReduceExecutable) {

}

type registryServer struct {
	pb.UnimplementedRegistryServer
	workers chan<- Worker

	executables []api.MapReduceExecutable
}

func (s *registryServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	missingExecutables := s.getMissingExecutables(in)
	if len(missingExecutables) > 0 {
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

func serveRegistry(lis net.Listener, s *grpc.Server) <-chan Worker {
	workers := make(chan Worker, WorkerLimit)
	go func() {
		defer close(workers)
		pb.RegisterRegistryServer(s, &registryServer{workers: workers})
		log.Printf("registryServer listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve registryServer: %v", err)
		}
	}()
	return workers
}

func (s* registryServer) getMissingExecutables(in *pb.RegisterRequest) []int {
	missing := make([]int, 0)

	Outer:
	for i := range s.executables {
		for _, m := range s.executables[i].GetMappers(){
			if ! contains(m, in.GetMappers()){
				missing = append(missing, i)
				continue Outer
			}
		}
		for _, r := range s.executables[i].GetReducers(){
			if ! contains(r, in.GetReducers()){
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
