package master

import (
	"fmt"
	"github.com/vynaloze/mapreduce/engine/master/controller"
	"time"
)

func Run(addr string) {
	c := controller.New(addr)

	for {
		fmt.Println(c.FreeWorkers(2))
		time.Sleep(time.Second)
	}
}
