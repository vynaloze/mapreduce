package master

import (
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	"github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/master/controller"
	"log"
)

type scheduler struct {
	controller   controller.Controller
	notifyClient chan<- *external.JobStatus

	mapTaskResults []*api.MapTaskStatus
}

func (s *scheduler) handleJob(job *external.Job) {
	log.Printf("received job: %v", job)

	s.notify("start partitioning input data")
	splits, err := s.split(job.GetSpec().GetIn())
	if err != nil {
		s.notify("split error: " + err.Error())
		return
	}
	s.notify(fmt.Sprintf("produced %d splits", len(splits)))

	s.notify("start map phase")
	s.mapPhase(splits, job.GetSpec().GetOut().GetOutputPartitions())
	s.notify(fmt.Sprintf("end map phase: received %d regions", len(s.mapTaskResults)))

	// TODO start with reduce tasks and take all remaining workers for map tasks (also later TODO)

}

func (s *scheduler) notify(msg string) {
	s.notifyClient <- &external.JobStatus{Message: msg}
}
