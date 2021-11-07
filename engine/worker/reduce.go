package worker

import (
	"context"
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	mrio "github.com/vynaloze/mapreduce/engine/io"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

type reduceWorkerServer struct {
	internal.UnimplementedReduceWorkerServer

	data map[string][]string
}

func (r *reduceWorkerServer) Reduce(task *internal.ReduceTask, stream internal.ReduceWorker_ReduceServer) error {
	log.Printf("received reduce task: %+v", task)
	spec := task.GetOutputSpec()
	switch format := spec.GetOutputFormat().GetFormat(); format {
	case external.FileFormat_CSV:
		filename := fmt.Sprintf("%s%04d.csv", spec.GetOutputLocation(), task.GetPartition())
		h := mrio.CsvHandler{Filename: filename}
		if err := r.reduce(&h); err != nil {
			return err
		}
		rts := internal.ReduceTaskStatus{
			Result: &external.DFSFile{
				Location: filename,
				Format:   spec.GetOutputFormat(),
			},
		}
		if err := stream.Send(&rts); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported output format: %+v", format)
	}
}

func (r *reduceWorkerServer) reduce(handler mrio.Handler) error {
	for k, vals := range r.data {
		err := r.reduceOne(handler, k, vals)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *reduceWorkerServer) reduceOne(handler mrio.Handler, k string, vals []string) error {
	wc := WordCount{}
	valsChan := make(chan *internal.Value, len(vals))
	go func() {
		for _, val := range vals {
			valsChan <- &internal.Value{Value: val}
		}
		close(valsChan)
	}()
	res := wc.Reduce(&internal.Key{Key: k}, valsChan)
	handler.Write(res)
	return nil
}

func (r *reduceWorkerServer) Notify(stream internal.ReduceWorker_NotifyServer) error {
	log.Printf("start downloading intermediate data")
	missingRegions := make([]*internal.Region, 0)
	for {
		region, err := stream.Recv()
		if err == io.EOF {
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
	conn, err := grpc.Dial(region.GetAddr(), grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect to map worker: %w", err)
	}
	c := internal.NewMapWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	stream, err := c.Get(ctx, region)
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

func (w *WordCount) Reduce(key *internal.Key, values <-chan *internal.Value) <-chan *internal.Pair {
	// key: a word
	// values: a list of counts
	o := make(chan *internal.Pair)
	go func() {
		defer close(o)

		var result int
		for v := range values {
			i, err := strconv.Atoi(v.GetValue())
			if err != nil {
				panic(err)
			}
			result += i
		}
		time.Sleep(100 * time.Millisecond) //FIXME

		o <- &internal.Pair{Key: key, Value: &internal.Value{Value: strconv.Itoa(result)}}
	}()
	return o
}
