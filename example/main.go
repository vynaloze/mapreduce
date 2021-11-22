package main

import (
	"github.com/vynaloze/mapreduce/api"
	"github.com/vynaloze/mapreduce/client"
)

func main() {
	c := client.New()

	c.RegisterMapper(client.MapperEntry{Name: "wordcount", Mapper: &WordCount{}})
	c.RegisterReducer(client.ReducerEntry{Name: "adder", Reducer: &Adder{}})


	format := &api.FileFormat{Format: api.FileFormat_TEXT}

	job := api.Job{
		Name: "example",
		Spec: &api.Spec{
			In: &api.InputSpec{
				InputFiles: []*api.DFSFile{
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/01.txt", Format: format, SizeBytes: 4 * 1024},
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/02.txt", Format: format, SizeBytes: 4 * 1024},
					//{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/03.txt", Format: format, SizeBytes: 4 * 1024}, FIXME error token too long in one of those files - TEST AND FIX IT
					//{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/04.txt", Format: format, SizeBytes: 4 * 1024},
				},
				InputSplitSizeBytes: 1.5 * 1024,
			},
			Out: &api.OutputSpec{
				OutputPartitions: 2,
				OutputLocation:   "/mnt/d/workspace/s2/3/mapreduce/example/output/",
				OutputFormat:     &api.FileFormat{Format: api.FileFormat_CSV},
			},
			Mapper: "wordcount",
			Reducer: "adder",
		},
	}

	c.SubmitAndWait(&job)
}
