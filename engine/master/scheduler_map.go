package master

import (
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/io"
	"log"
)

func (s *scheduler) split(spec *external.InputSpec) ([]internal.Split, error) {
	textHandler := &io.TextHandler{}

	splits := make([]internal.Split, 0)
	for _, file := range spec.GetInputFiles() {
		switch format := file.GetFormat().GetFormat(); format {
		case external.FileFormat_TEXT:
			splits = append(splits, textHandler.Split(spec)...)
		default:
			return nil, fmt.Errorf("unsupported format: %+v", format)
		}
	}
	return splits, nil
}

func (s *scheduler) mapPhase(splits []internal.Split) {
	mapTasks := make(chan *internal.MapTask, len(splits))
	mapTaskResults := make(chan *internal.MapTaskStatus)
	for i := range splits {
		mapTasks <- &internal.MapTask{InputSplit: &splits[i]}
	}
	close(mapTasks)
	go s.controller.ProcessMapTasks(mapTasks, mapTaskResults)
	for mtr := range mapTaskResults {
		log.Printf("received map task result: %+v", mtr)
		s.mapTaskResults = append(s.mapTaskResults, mtr)
		log.Printf("total results: %d", len(s.mapTaskResults))
	}
}
