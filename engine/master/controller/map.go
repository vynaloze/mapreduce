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

const concurrentMapTaskLimit = 5 //TODO dynamic

func (c *controller) ProcessMapTasks(tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus) {
	var wg sync.WaitGroup
	for i := 0; i < concurrentMapTaskLimit; i++ {
		wg.Add(1)
		log.Printf("started map processing #%d", i)
		go c.processMapTasks(&wg, tasks, results)
	}
	wg.Wait()
	close(results)
}

func (c *controller) processMapTasks(wg *sync.WaitGroup, tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus) {
	for task := range tasks {
		for {
			err := c.tryProcessMapTasks(task, results)
			if err != nil {
				log.Println(err)
			} else {
				break
			}
		}
	}
	wg.Done()
}

func (c *controller) tryProcessMapTasks(task *internal.MapTask, results chan<- *internal.MapTaskStatus) error {
	var w *MapWorker
	for {
		w = c.getNextFreeWorkerForMap()
		if w != nil {
			break
		}
		log.Println("no free workers for map task - try again in 1 sec")
		time.Sleep(time.Second)
	}
	defer c.freeMapWorker(w)
	defer w.conn.Close()

	stream, err := w.client.Map(context.Background(), task)
	if err != nil {
		return fmt.Errorf("error: request w.Map(%+v): %w\n", task, err)

	}
	for {
		mts, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error: streaming w.Map(%+v): %w\n", task, err)
		}
		results <- mts
	}
	return nil
}

type MapWorker struct {
	*Worker
	conn   *grpc.ClientConn
	client internal.MapWorkerClient
}

func NewMapWorker(w *Worker) *MapWorker {
	conn, err := grpc.Dial(w.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	c := internal.NewMapWorkerClient(conn)
	return &MapWorker{
		Worker: w,
		conn:   conn,
		client: c,
	}
}
