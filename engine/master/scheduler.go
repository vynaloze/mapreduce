package master

import (
	"fmt"
	"github.com/google/uuid"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/master/controller"
	"log"
	"sync"
)

type scheduler struct {
	controller   controller.Controller
	notifyClient chan<- *external.JobStatus
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
	maxRegions := int64(len(splits)) * job.GetSpec().GetOut().GetOutputPartitions()

	s.notify("start map phase")
	mapTasks := make(map[string]*internal.MapTask)
	initialMapTasksChan := make(chan *internal.MapTask, len(splits))
	for i := range splits {
		mt := &internal.MapTask{Id: uuid.New().String(), InputSplit: &splits[i], Partitions: job.GetSpec().GetOut().GetOutputPartitions()}
		mapTasks[mt.GetId()] = mt
		initialMapTasksChan <- mt
	}
	close(initialMapTasksChan)
	mapTaskResults := make(map[int64][]*internal.MapTaskStatus, 0)
	initialMapTaskResultsChan := make(chan *internal.MapTaskStatus)
	go s.controller.ProcessMapTasks(initialMapTasksChan, initialMapTaskResultsChan)

	s.notify("start reduce phase")
	reduceTasks := make(map[int64]*controller.ReduceTask, job.GetSpec().GetOut().GetOutputPartitions())
	reduceTasksChan := make(chan *controller.ReduceTask, job.GetSpec().GetOut().GetOutputPartitions())
	reduceRegions := make(map[int64]chan *internal.Region, job.GetSpec().GetOut().GetOutputPartitions())
	rerunReduce := make(chan int64)
	for i := int64(0); i < job.GetSpec().GetOut().GetOutputPartitions(); i++ {
		rt := &controller.ReduceTask{
			Task:        &internal.ReduceTask{OutputSpec: job.GetSpec().GetOut(), Partition: i},
			Regions:     make(chan chan *internal.Region, 1),
			RerunReduce: rerunReduce,
		}
		reduceTasks[i] = rt
		reduceTasksChan <- rt

		rr := make(chan *internal.Region, maxRegions)
		reduceRegions[i] = rr
		rt.Regions <- rr
	}
	close(reduceTasksChan)
	rerunMapTasks := make(chan string)
	reduceTaskResultsChan := make(chan *internal.ReduceTaskStatus)
	go s.controller.ProcessReduceTasks(reduceTasksChan, rerunMapTasks, reduceTaskResultsChan)

	rerunMapTaskResultsChan := make(chan *internal.MapTaskStatus)

	mux := sync.Mutex{}
	go func() {
		for rerunMapTaskId := range rerunMapTasks {
			mux.Lock()
			log.Printf("map worker failure during reduce phase: reruning task %+v", mapTasks[rerunMapTaskId])
			go s.controller.ProcessMapTask(mapTasks[rerunMapTaskId], rerunMapTaskResultsChan)
			mux.Unlock()
		}
	}()
	go func() {
		for rerunReduceTaskPartition := range rerunReduce {
			mux.Lock()
			log.Printf("reduce task %d failure: request all regions again", rerunReduceTaskPartition)
			r := make(chan *internal.Region, maxRegions)
			reduceRegions[rerunReduceTaskPartition] = r
			reduceTasks[rerunReduceTaskPartition].Regions <- r
			for _, mtr := range mapTaskResults[rerunReduceTaskPartition] {
				r <- mtr.GetRegion()
			}
			// TODO close??? where???
			mux.Unlock()
		}
	}()
	go func() {
		for mtr := range initialMapTaskResultsChan {
			mux.Lock()
			log.Printf("received initial map task result: %+v", mtr)
			p := mtr.GetRegion().GetPartition()
			_, ok := mapTaskResults[p]
			if !ok {
				mapTaskResults[p] = make([]*internal.MapTaskStatus, 0)
			}
			mapTaskResults[p] = append(mapTaskResults[p], mtr)
			reduceRegions[p] <- mtr.GetRegion()
			mux.Unlock()
		}
		log.Println("finished all initial map tasks")
		for _, c := range reduceRegions {
			close(c)
		}
	}()
	go func() {
		for mtr := range rerunMapTaskResultsChan {
			mux.Lock()
			log.Printf("received rerun map task result: %+v", mtr)
			p := mtr.GetRegion().GetPartition()
			_, ok := mapTaskResults[p]
			if !ok {
				mapTaskResults[p] = make([]*internal.MapTaskStatus, 0)
			}
			mapTaskResults[p] = append(mapTaskResults[p], mtr)
			reduceRegions[p] <- mtr.GetRegion()
			mux.Unlock()
		}
	}()

	for rts := range reduceTaskResultsChan {
		log.Printf("received reduce task result: %+v", rts)
	}
	s.notify("finished mapreduce")
}

func (s *scheduler) notify(msg string) {
	s.notifyClient <- &external.JobStatus{Message: msg}
}
