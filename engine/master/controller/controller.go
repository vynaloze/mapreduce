package controller

import (
	internal "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const (
	WorkerLimit = 3
	// right now, this is just used as a limit of concurrent Register calls (buffer of chan)
	// TODO enforce this limit or rename it
)

type Worker struct {
	Uuid      string
	Addr      string
	ExpiresOn time.Time
}

type Controller interface {
	ProcessMapTasks(tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus)
}

type controller struct {
	workers    map[string]Worker
	mapWorkers map[string]MapWorker
	//reduceWorkers    map[string]Worker
	workersMux sync.RWMutex

	newWorkers <-chan string
}

func New(lis net.Listener, s *grpc.Server) Controller {
	c := &controller{
		workers:    make(map[string]Worker, 0),
		mapWorkers: make(map[string]MapWorker, 0),
	}
	go c.receiveNewWorkers(serveRegistry(lis, s))
	return c
}

func (c *controller) getNextFreeWorkerForMap() *MapWorker {
	c.workersMux.Lock()
	defer c.workersMux.Unlock()

	for uuid := range c.workers {
		notExpired := time.Now().Before(c.workers[uuid].ExpiresOn)
		_, isMapWorker := c.mapWorkers[uuid]
		if notExpired && !isMapWorker { //TODO isReduceWorker
			w := c.workers[uuid]
			cw := NewMapWorker(&w)
			c.mapWorkers[uuid] = *cw
			return cw
		}
	}
	return nil
}

func (c *controller) freeMapWorker(w *MapWorker) {
	c.workersMux.Lock()
	defer c.workersMux.Unlock()
	delete(c.mapWorkers, w.Uuid)
}

func (c *controller) receiveNewWorkers(newWorkers <-chan Worker) {
	for {
		select {
		case w := <-newWorkers:
			c.workersMux.Lock()
			c.workers[w.Uuid] = w
			c.workersMux.Unlock()
		}
	}
}
