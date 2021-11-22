package public

import (
	"context"
	external "github.com/vynaloze/mapreduce/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewMapReduceRegistry(lis net.Listener, s *grpc.Server) <-chan *external.MapReduceExecutable {
	return serveMRR(lis, s)
}

type mapreduceRegistry struct {
	external.UnimplementedMapReduceRegistryServer
	executables chan *external.MapReduceExecutable
}

func (r * mapreduceRegistry) Register(c context.Context, e *external.MapReduceExecutable) (*external.Empty, error) {
	r.executables <- e
	return &external.Empty{}, nil
}

func serveMRR(lis net.Listener, s *grpc.Server) <-chan *external.MapReduceExecutable {
	e := make(chan *external.MapReduceExecutable)

	go func() {
		defer close(e)

		external.RegisterMapReduceRegistryServer(s, &mapreduceRegistry{executables: e})
		log.Printf("mapreduceRegistry listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve mapreduceRegistry: %v", err)
		}
	}()
	return e
}
