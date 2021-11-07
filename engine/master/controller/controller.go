package controller

import (
	internal "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
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
	Expired   bool
}

type Controller interface {
	ProcessMapTasks(tasks <-chan *internal.MapTask, results chan<- *internal.MapTaskStatus)
	ProcessMapTask(task *internal.MapTask, results chan<- *internal.MapTaskStatus)
	ProcessReduceTasks(tasks <-chan *ReduceTask, rerunMapTasks chan<- RerunMapTasks, results chan<- *internal.ReduceTaskStatus)
}

type controller struct {
	workers       map[string]Worker
	mapWorkers    map[string]MapWorker
	reduceWorkers map[string]ReduceWorker
	workersMux    sync.RWMutex

	newWorkers <-chan string
}

func New(lis net.Listener, s *grpc.Server) Controller {
	c := &controller{
		workers:       make(map[string]Worker),
		mapWorkers:    make(map[string]MapWorker),
		reduceWorkers: make(map[string]ReduceWorker),
	}
	go c.receiveNewWorkers(serveRegistry(lis, s))
	go c.workerStats()
	return c
}

func (c *controller) getNextFreeWorkerForMap() *MapWorker {
	c.workersMux.Lock()
	defer c.workersMux.Unlock()

	for uuid := range c.workers {
		notExpired := time.Now().Before(c.workers[uuid].ExpiresOn)
		_, isMapWorker := c.mapWorkers[uuid]
		_, isReduceWorker := c.reduceWorkers[uuid]
		if notExpired && !c.workers[uuid].Expired && !isMapWorker && !isReduceWorker {
			w := c.workers[uuid]
			cw, err := NewMapWorker(&w)
			if err != nil {
				log.Println(err)
				w.Expired = true
				c.workers[uuid] = w
				continue
			}
			c.mapWorkers[uuid] = *cw
			return cw
		}
	}
	return nil
}

func (c *controller) getNextFreeWorkerForReduce() *ReduceWorker {
	c.workersMux.Lock()
	defer c.workersMux.Unlock()

	for uuid := range c.workers {
		notExpired := time.Now().Before(c.workers[uuid].ExpiresOn)
		_, isMapWorker := c.mapWorkers[uuid]
		_, isReduceWorker := c.reduceWorkers[uuid]
		if notExpired && !c.workers[uuid].Expired && !isMapWorker && !isReduceWorker {
			w := c.workers[uuid]
			cw, err := NewReduceWorker(&w)
			if err != nil {
				log.Println(err)
				w.Expired = true
				c.workers[uuid] = w
				continue
			}
			c.reduceWorkers[uuid] = *cw
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

func (c *controller) freeReduceWorker(w *ReduceWorker) {
	c.workersMux.Lock()
	defer c.workersMux.Unlock()
	delete(c.reduceWorkers, w.Uuid)
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

func (c *controller) workerStats() string {
	for {
		var free, mapper, reducer, dead int
		c.workersMux.RLock()
		for uuid := range c.workers {
			expired := !time.Now().Before(c.workers[uuid].ExpiresOn) || c.workers[uuid].Expired
			_, isMapWorker := c.mapWorkers[uuid]
			_, isReduceWorker := c.reduceWorkers[uuid]
			switch {
			case expired:
				dead++
			case isMapWorker:
				mapper++
			case isReduceWorker:
				reducer++
			default:
				free++
			}
		}
		c.workersMux.RUnlock()
		log.Printf("free=%d,mapper=%d,reducer=%d,dead=%d", free, mapper, reducer, dead)
		time.Sleep(5 * time.Second)
	}
}
