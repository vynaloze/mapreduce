package worker

import (
	"context"
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	mrio "github.com/vynaloze/mapreduce/engine/io"
	"hash/fnv"
	"io"
	"log"
	"sync"
	"time"
)

const bufferSizePerPartition = 1000

type mapWorkerServer struct {
	internal.UnimplementedMapWorkerServer
	addr string

	taskPartitions map[string]*taskPartitions
	mux            sync.RWMutex
}

type taskPartitions struct {
	partitions map[int64]chan *external.Pair
	mux        sync.RWMutex
}

func (s *mapWorkerServer) Map(task *internal.MapTask, stream internal.MapWorker_MapServer) error {
	log.Printf("received map task: %+v", task)

	s.mux.Lock()
	s.taskPartitions[task.GetId()] = &taskPartitions{partitions: make(map[int64]chan *external.Pair)}
	s.mux.Unlock()
	defer s.cleanup(task.GetId())

	th := &mrio.TextHandler{}

	for mrs == nil || mrs.client == nil {
		log.Printf("waiting for map function to be ready")
		time.Sleep(1 * time.Second)
	}

	for inputPair := range th.Read(task.GetInputSplit()) {
		intermediatePairs := mrs.Map(inputPair.GetKey(), inputPair.GetValue())
		for pair := range intermediatePairs {
			partition := hash(pair.GetKey().GetKey(), task.GetPartitions())

			s.mux.RLock()
			partitions := s.taskPartitions[task.GetId()]
			s.mux.RUnlock()

			partitions.mux.RLock()
			_, ok := partitions.partitions[partition]
			partitions.mux.RUnlock()
			if !ok {
				// new partition
				c := make(chan *external.Pair, bufferSizePerPartition)
				//defer close(c)
				partitions.mux.Lock()
				partitions.partitions[partition] = c
				partitions.mux.Unlock()

				mts := internal.MapTaskStatus{
					Task: task,
					Region: &internal.Region{
						Addr:      s.addr,
						Partition: partition,
						TaskId:    task.GetId(),
					},
				}
				if err := stream.Send(&mts); err != nil {
					return err
				}
			}
			partitions.mux.RLock()
			partitions.partitions[partition] <- pair
			partitions.mux.RUnlock()
		}
	}
	return nil
}

func (s *mapWorkerServer) cleanup(taskId string) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	s.taskPartitions[taskId].cleanup()
}

func (tp *taskPartitions) cleanup() {
	for _, c := range tp.partitions {
		close(c)
	}
}

func (s *mapWorkerServer) Get(region *internal.Region, stream internal.MapWorker_GetServer) error {
	s.mux.RLock()
	ps, ok := s.taskPartitions[region.GetTaskId()]
	if !ok {
		return fmt.Errorf("no result for task %s", region.GetTaskId())
	}
	s.mux.RUnlock()
	ps.mux.RLock()
	p, ok := ps.partitions[region.GetPartition()]
	if !ok {
		return fmt.Errorf("no result for partition %d", region.GetPartition())
	}
	ps.mux.RUnlock()

	for pair := range p {
		if err := stream.Send(pair); err != nil {
			return err
		}
	}
	return nil
}

func hash(s string, r int64) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64() % uint64(r))
}

func (mr *mapReduceServer) Map(key *external.Key, value *external.Value) <-chan *external.Pair {
	o := make(chan *external.Pair)
	go func() {
		defer close(o)
		stream, err := mr.client.Map(context.TODO(), &external.Pair{Key: key, Value: value})

		if err != nil {
			log.Fatalf("error calling internal Map: %s", err)
		}
		for {
			pair, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.Map(_) = _, %v", mr.client, err)
			}
			o <- pair
		}
	}()
	return o
}
