package master

import (
	"fmt"
	"sync"
	"time"
)

func Run(addr string) {
	workersMux := sync.RWMutex{}
	workers := make(map[string]struct{}, 0)
	registrar := make(chan string)

	go serveRegistrar(addr, registrar)

	for {
		select {
		case w := <-registrar:
			workersMux.Lock()
			workers[w] = struct{}{}
			workersMux.Unlock()
		default:
			fmt.Println(workers)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
