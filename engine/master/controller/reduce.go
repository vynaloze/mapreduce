package controller

import (
	"context"
	"fmt"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

const concurrentReduceTaskLimit = 2 //TODO dynamic

type ReduceTask struct {
	Task        *internal.ReduceTask
	Regions     chan []*internal.Region
	RerunReduce chan int64
}

type RerunMapTasks struct {
	MapTaskIds []string
	Partition  int64
}

func (c *controller) ProcessReduceTasks(tasks <-chan *ReduceTask, rerunMapTasks chan<- RerunMapTasks, results chan<- *internal.ReduceTaskStatus) {
	var wg sync.WaitGroup
	for i := 0; i < concurrentReduceTaskLimit; i++ {
		wg.Add(1)
		log.Printf("started reduce processing #%d", i)
		go c.processReduceTasks(&wg, tasks, rerunMapTasks, results)
	}
	wg.Wait()
	log.Printf("finished reduce processing")
	close(rerunMapTasks)
	close(results)
}

func (c *controller) processReduceTasks(wg *sync.WaitGroup, tasks <-chan *ReduceTask, rerunMapTasks chan<- RerunMapTasks, results chan<- *internal.ReduceTaskStatus) {
	for task := range tasks {
		for {
			err := c.tryProcessReduceTask(task, rerunMapTasks, results)
			if err != nil {
				log.Printf("error during reduce: %s", err)
				task.RerunReduce <- task.Task.GetPartition()
			} else {
				//close(task.Regions)
				break
			}
		}
	}
	wg.Done()
}

func (c *controller) tryProcessReduceTask(task *ReduceTask, rerunMapTasks chan<- RerunMapTasks, results chan<- *internal.ReduceTaskStatus) error {
	var w *ReduceWorker
	for {
		w = c.getNextFreeWorkerForReduce()
		if w != nil {
			break
		}
		log.Println("no free workers for reduce task - try again in 5 sec")
		time.Sleep(5 * time.Second)
	}
	defer c.freeReduceWorker(w)
	defer w.conn.Close()

	// ensure the worker has all intermediate data (from map workers)
	for regions := range task.Regions {
		missingRegions, err := w.notifyOnce(regions)
		if err != nil {
			return err
		}
		if len(missingRegions.GetRegions()) == 0 {
			break
		} else {
			taskIds := make([]string, 0)
			for _, region := range missingRegions.GetRegions() {
				taskIds = append(taskIds, region.GetTaskId())
			}
			rerunMapTasks <- RerunMapTasks{MapTaskIds: taskIds, Partition: task.Task.GetPartition()}
		}
	}

	log.Printf("start reduce task #%d", task.Task.GetPartition())
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	stream, err := w.client.Reduce(ctx, task.Task)
	if err != nil {
		return fmt.Errorf("error: request w.Reduce(%+v): %w\n", task.Task, err)
	}
	for {
		rts, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: streaming w.Reduce(%+v): %w\n", task.Task, err)
		}
		results <- rts
	}
	return nil
}

type ReduceWorker struct {
	*Worker
	conn   *grpc.ClientConn
	client internal.ReduceWorkerClient
}

func NewReduceWorker(w *Worker) (*ReduceWorker, error) {
	conn, err := grpc.Dial(w.Addr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not establish connection to reduce worker: %w", err)
	}
	c := internal.NewReduceWorkerClient(conn)
	return &ReduceWorker{
		Worker: w,
		conn:   conn,
		client: c,
	}, nil
}

func (rw *ReduceWorker) notifyOnce(regions []*internal.Region) (*internal.MissingRegions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	stream, err := rw.client.Notify(ctx)
	if err != nil {
		return nil, fmt.Errorf("error: request w.Notify(<Regions>): %w\n", err)
	}
	for _, region := range regions {
		if region == nil {
			log.Fatalf("Notify(): region==nil") //FIXME
		}
		if err := stream.Send(region); err != nil {
			return nil, fmt.Errorf("%v.Notify(%v) = %v", stream, region, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("%v.CloseAndRecv() got error %v", stream, err)
	}
	return reply, nil
}
