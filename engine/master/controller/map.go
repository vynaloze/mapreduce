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

const concurrentMapTaskLimit = 3 //TODO dynamic

func (c *controller) ProcessMapTasks(tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus) {
	defer close(results)
	var wg sync.WaitGroup
	for i := 0; i < concurrentMapTaskLimit; i++ {
		wg.Add(1)
		log.Printf("started map processing #%d", i)
		go c.processMapTasks(&wg, tasks, results)
	}
	wg.Wait()
}

func (c *controller) processMapTasks(wg *sync.WaitGroup, tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus) {
	for task := range tasks {
		c.ProcessMapTask(task, results)
	}
	wg.Done()
}

func (c *controller) ProcessMapTask(task *internal.MapTask, results chan<- *internal.MapTaskStatus) {
	for {
		err := c.tryProcessMapTask(task, results)
		if err != nil {
			log.Printf("error during map: %s", err)
		} else {
			break
		}
	}
}

func (c *controller) tryProcessMapTask(task *internal.MapTask, results chan<- *internal.MapTaskStatus) error {
	var w *MapWorker
	for {
		w = c.getNextFreeWorkerForMap()
		if w != nil {
			break
		}
		log.Println("no free workers for map task - try again in 5 sec")
		time.Sleep(5 * time.Second)
	}
	defer c.freeMapWorker(w)
	defer w.conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := w.client.Map(ctx, task)
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

func NewMapWorker(w *Worker) (*MapWorker, error) {
	conn, err := grpc.Dial(w.Addr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not establish connection to map worker: %w", err)
	}
	c := internal.NewMapWorkerClient(conn)
	return &MapWorker{
		Worker: w,
		conn:   conn,
		client: c,
	}, nil
}
