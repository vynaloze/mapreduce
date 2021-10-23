package controller

import (
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
	Task      bool //TODO
}

type Controller interface {
	FreeWorkers(count int) []Worker
}

type controller struct {
	workers    map[string]Worker
	workersMux sync.RWMutex

	newWorkers <-chan string
}

func New(lis net.Listener, s *grpc.Server) Controller {
	c := &controller{
		workers: make(map[string]Worker, 0),
	}
	go c.receiveNewWorkers(serveRegistry(lis, s))
	return c
}

func (c *controller) FreeWorkers(count int) []Worker {
	workers := make([]Worker, 0)

	c.workersMux.RLock()
	defer c.workersMux.RUnlock()

	i := 0
	for _, w := range c.workers {
		if count != -1 && i >= count {
			break
		}
		if time.Now().Before(w.ExpiresOn) && w.Task == false {
			workers = append(workers, w)
		}
		i++
	}
	return workers
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
