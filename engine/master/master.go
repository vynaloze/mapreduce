package master

import (
	"fmt"
	"github.com/vynaloze/mapreduce/api"
	"github.com/vynaloze/mapreduce/engine/master/controller"
	"github.com/vynaloze/mapreduce/engine/master/public"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func Run(addr string) {
	lis, s := initServer(addr)

	c := controller.New(lis, s)
	jobs, jobStatus := public.New(lis, s)

	for {
		select {
		case j := <-jobs:
			log.Printf("received job: %v", j)
			time.Sleep(time.Second)
			jobStatus <- &api.JobStatus{Message: "ack: " + time.Now().String()}
			time.Sleep(2 * time.Second)
			jobStatus <- &api.JobStatus{Message: "done: " + time.Now().String()}
			close(jobStatus)
		default:
			fmt.Println(c.FreeWorkers(2))
			time.Sleep(time.Second)
		}
	}
}

func initServer(addr string) (net.Listener, *grpc.Server) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	return lis, s
}
