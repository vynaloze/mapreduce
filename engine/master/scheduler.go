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

	s.notify("start map phase")
	mapTasks := make(map[string]*internal.MapTask)
	initialMapTasksChan := make(chan *internal.MapTask, len(splits))
	for i := range splits {
		mt := &internal.MapTask{
			Id: uuid.New().String(),
			InputSplit: &splits[i],
			Partitions: job.GetSpec().GetOut().GetOutputPartitions(),
			Mapper: job.GetSpec().GetMapper(),
		}
		mapTasks[mt.GetId()] = mt
		initialMapTasksChan <- mt
	}
	close(initialMapTasksChan)
	mapTaskResults := make(map[int64][]*internal.MapTaskStatus, 0)
	mapTaskResultsBatches := make(chan chan *internal.MapTaskStatus, 1)
	initialBatch := make(chan *internal.MapTaskStatus)
	mapTaskResultsBatches <- initialBatch
	go s.controller.ProcessMapTasks(initialMapTasksChan, initialBatch)

	s.notify("prepare reduce phase")
	reduceTasks := make(map[int64]*controller.ReduceTask, job.GetSpec().GetOut().GetOutputPartitions())
	reduceTasksChan := make(chan *controller.ReduceTask, job.GetSpec().GetOut().GetOutputPartitions())
	reduceRegions := make(map[int64][]*internal.Region, job.GetSpec().GetOut().GetOutputPartitions())
	rerunReduce := make(chan int64)
	for i := int64(0); i < job.GetSpec().GetOut().GetOutputPartitions(); i++ {
		rt := &controller.ReduceTask{
			Task:        &internal.ReduceTask{OutputSpec: job.GetSpec().GetOut(), Partition: i, Reducer: job.GetSpec().GetMapper()},
			Regions:     make(chan []*internal.Region, 1),
			RerunReduce: rerunReduce,
		}
		reduceTasks[i] = rt
		reduceTasksChan <- rt

		//rr := make([]*internal.Region, maxRegions)
		//reduceRegions[i] = rr
		//rt.Regions <- rr
	}
	rerunMapTasks := make(chan controller.RerunMapTasks)
	reduceTaskResultsChan := make(chan *internal.ReduceTaskStatus)
	go s.controller.ProcessReduceTasks(reduceTasksChan, rerunMapTasks, reduceTaskResultsChan)

	mux := sync.Mutex{}
	go func() {
		for rerunMapTask := range rerunMapTasks {
			mux.Lock()
			log.Printf("map worker failure during reduce phase: reruning tasks %+v", rerunMapTask)
			// clear results with failed task
			for _, results := range mapTaskResults {
				okResults := make([]*internal.MapTaskStatus, 0)
				for _, r := range results {
					for _, rrid := range rerunMapTask.MapTaskIds {
						if r.GetTask().GetId() != rrid || r.GetRegion().GetPartition() != rerunMapTask.Partition {
							okResults = append(okResults, r)
						}
					}
				}
				results = okResults
			}
			log.Printf("cleared results containing failed tasks %+v", rerunMapTask)

			// notify about new regions to download
			//reduceRegions[rerunMapTask.Partition] = make([]*internal.Region, maxRegions)

			// rerun
			tasks := make(chan *internal.MapTask, len(rerunMapTask.MapTaskIds))
			for _, rrid := range rerunMapTask.MapTaskIds {
				tasks <- mapTasks[rrid]
				log.Printf("resubmitted map task %s", rrid)
			}
			close(tasks)
			resBatch := make(chan *internal.MapTaskStatus)
			mapTaskResultsBatches <- resBatch
			go s.controller.ProcessMapTasks(tasks, resBatch)
			mux.Unlock()
		}
	}()
	go func() {
		for rerunReduceTaskPartition := range rerunReduce {
			mux.Lock()
			log.Printf("reduce task #%d failure: request all regions again", rerunReduceTaskPartition)
			r := make([]*internal.Region, 0)
			for _, mtr := range mapTaskResults[rerunReduceTaskPartition] {
				r = append(r, mtr.GetRegion())
			}
			reduceRegions[rerunReduceTaskPartition] = r
			reduceTasks[rerunReduceTaskPartition].Regions <- r
			log.Printf("submitted region replay for reduce task #%d", rerunReduceTaskPartition)
			mux.Unlock()
		}
	}()
	go func() {
		for mapTaskBatch := range mapTaskResultsBatches {
			var tasksNo int

			for i := range reduceRegions {
				reduceRegions[i] = make([]*internal.Region, 0)
			}
			for mtr := range mapTaskBatch {
				tasksNo++
				mux.Lock()
				log.Printf("received initial map task result: %+v", mtr.GetRegion())
				p := mtr.GetRegion().GetPartition()
				_, ok := mapTaskResults[p]
				if !ok {
					mapTaskResults[p] = make([]*internal.MapTaskStatus, 0)
				}
				mapTaskResults[p] = append(mapTaskResults[p], mtr)
				reduceRegions[p] = append(reduceRegions[p], mtr.GetRegion())
				mux.Unlock()
			}
			s.notify(fmt.Sprintf("finished all (%d) map tasks in current batch", tasksNo))
			s.notify("start reduce phase")
			for p, rt := range reduceTasks {
				rt.Regions <- reduceRegions[p]
			}
		}
	}()

	var finishedReduceTasks int64
	for rts := range reduceTaskResultsChan {
		log.Printf("received reduce task result: %+v", rts)
		finishedReduceTasks++
		if finishedReduceTasks >= job.GetSpec().GetOut().GetOutputPartitions() {
			s.notify(fmt.Sprintf("finished all %d reduce tasks", finishedReduceTasks))
			close(reduceTasksChan)
		}
	}
	s.notify("finished job")
}

func (s *scheduler) notify(msg string) {
	log.Println(msg)
	s.notifyClient <- &external.JobStatus{Message: msg}
}
