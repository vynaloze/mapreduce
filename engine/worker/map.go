package worker

import (
	"fmt"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/io"
	"hash/fnv"
	"log"
	"strings"
	"sync"
	"time"
)

const bufferSizePerPartition = 1000

type mapWorkerServer struct {
	pb.UnimplementedMapWorkerServer
	addr string

	taskPartitions map[string]*taskPartitions
	mux            sync.RWMutex
}

type taskPartitions struct {
	partitions map[int64]chan *pb.Pair
	mux        sync.RWMutex
}

func (s *mapWorkerServer) Map(task *pb.MapTask, stream pb.MapWorker_MapServer) error {
	log.Printf("received map task: %+v", task)

	s.mux.Lock()
	s.taskPartitions[task.GetId()] = &taskPartitions{partitions: make(map[int64]chan *pb.Pair)}
	s.mux.Unlock()
	defer s.cleanup(task.GetId())

	wc := &WordCount{}
	th := &io.TextHandler{}

	for inputPair := range th.Read(task.GetInputSplit()) {
		intermediatePairs := wc.Map(inputPair.GetKey(), inputPair.GetValue())
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
				c := make(chan *pb.Pair, bufferSizePerPartition)
				//defer close(c)
				partitions.mux.Lock()
				partitions.partitions[partition] = c
				partitions.mux.Unlock()

				mts := pb.MapTaskStatus{
					Task: task,
					Region: &pb.Region{
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

func (s *mapWorkerServer) Get(region *pb.Region, stream pb.MapWorker_GetServer) error {
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

type WordCount struct{}

func (w *WordCount) Map(key *pb.Key, value *pb.Value) <-chan *pb.Pair {
	// key: document name
	// value: document contents
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)

		for _, word := range strings.Fields(value.GetValue()) {
			k := strings.ReplaceAll(word, ",", "")
			kk := strings.ReplaceAll(k, ".", "")

			o <- &pb.Pair{Key: &pb.Key{Key: kk}, Value: &pb.Value{Value: "1"}}
		}
		time.Sleep(2000 * time.Millisecond) // FIXME
	}()
	return o
}
