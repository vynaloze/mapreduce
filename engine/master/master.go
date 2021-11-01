package master

import (
	"github.com/vynaloze/mapreduce/engine/api"
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

	scheduler := scheduler{
		c,
		jobStatus,
		make([]*api.MapTaskStatus, 0),
	}

	for {
		select {
		case j := <-jobs:
			scheduler.handleJob(j)
		default:
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
