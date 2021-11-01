package worker

import (
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

type mapWorkerServer struct {
	pb.UnimplementedMapWorkerServer
	addr string
}

func (s *mapWorkerServer) Map(task *pb.MapTask, stream pb.MapWorker_MapServer) error {
	log.Printf("received task: %+v", task)
	for i := 0; i < 2; i++ {
		mts := pb.MapTaskStatus{
			Task: task,
			Region: &pb.Region{
				Addr:      s.addr,
				Partition: int64(i),
			},
		}
		if err := stream.Send(&mts); err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func (s *mapWorkerServer) Get(region *pb.Region, stream pb.MapWorker_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// TODO actual Map:

type WordCount struct{}

func (w *WordCount) Map(key, value string) <-chan *pb.Pair {
	// key: document name
	// value: document contents
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)

		for _, word := range strings.Fields(value) {
			o <- &pb.Pair{Key: &pb.Key{Key: word}, Value: &pb.Value{Value: "1"}}
		}
	}()
	return o
}
