package worker

import (
	"context"
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type reduceWorkerServer struct {
	internal.UnimplementedReduceWorkerServer

	data map[string][]string
}

func (r *reduceWorkerServer) Reduce(task *internal.ReduceTask, stream internal.ReduceWorker_ReduceServer) error {
	log.Printf("received reduce task: %+v", task)
	for i := 0; i < 2; i++ {
		rts := internal.ReduceTaskStatus{
			Result: &external.DFSFile{
				Location:  fmt.Sprintf("foo_%d", i),
				Format:    &external.FileFormat{Format: external.FileFormat_TEXT},
				SizeBytes: 0,
			},
		}
		if err := stream.Send(&rts); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (r *reduceWorkerServer) Notify(stream internal.ReduceWorker_NotifyServer) error {
	log.Printf("notify")
	missingRegions := make([]*internal.Region, 0)
	for {
		region, err := stream.Recv()
		if err == io.EOF {
			log.Printf("data: %v", r.data)
			return stream.SendAndClose(&internal.MissingRegions{Regions: missingRegions})
		}
		if err != nil {
			return err
		}
		err = r.getFromMapWorker(region)
		if err != nil {
			log.Printf("cannot get region: %s", err)
			missingRegions = append(missingRegions, region)
		}
	}
}

func (r *reduceWorkerServer) getFromMapWorker(region *internal.Region) error {
	conn, err := grpc.Dial(region.GetAddr(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect to map worker: %w", err)
	}
	c := internal.NewMapWorkerClient(conn)
	stream, err := c.Get(context.Background(), region)
	if err != nil {
		return fmt.Errorf("error: request Get(%+v): %w\n", region, err)

	}
	for {
		pair, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: streaming Get(%+v): %w\n", region, err)
		}
		_, ok := r.data[pair.GetKey().GetKey()]
		if !ok {
			r.data[pair.GetKey().GetKey()] = make([]string, 0)
		}
		r.data[pair.GetKey().GetKey()] = append(r.data[pair.GetKey().GetKey()], pair.GetValue().GetValue())
	}
	return nil
}

//func (w *WordCount) Reduce(key string, values[]string) <-chan string {
//	// key: a word
//	// values: a list of counts
//	o := make(chan string)
//	go func() {
//		defer close(o)
//
//		var result int
//		for v := range values {
//			i, err := strconv.Atoi(v)
//			if err != nil {
//				panic(err)
//			}
//			result += i
//		}
//
//		o <- strconv.Itoa(result)
//	}()
//	return o
//}
