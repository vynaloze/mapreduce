package main

import (
	"context"
	"github.com/vynaloze/mapreduce/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	format := &api.FileFormat{Format: api.FileFormat_TEXT}

	job := api.Job{
		Name: "example",
		Spec: &api.Spec{
			In: &api.InputSpec{
				InputFiles: []*api.DFSFile{
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/01.txt", Format: format, SizeBytes: 4 * 1024},
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/02.txt", Format: format, SizeBytes: 4 * 1024},
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/03.txt", Format: format, SizeBytes: 4 * 1024},
					{Location: "/mnt/d/workspace/s2/3/mapreduce/example/input/04.txt", Format: format, SizeBytes: 4 * 1024},
				},
				InputSplitSizeBytes: 1.5 * 1024,
			},
			Out: &api.OutputSpec{
				OutputPartitions: 2,
				OutputLocation:   "/mnt/d/workspace/s2/3/mapreduce/example/output/",
				OutputFormat:     format,
			},
		},
	}

	// TODO this boilerplate should go to client package
	conn, err := grpc.Dial(":50050", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewMasterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour) //TODO correct timeout
	defer cancel()
	stream, err := c.Submit(ctx, &job)
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}

	for {
		status, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.Submit(_) = _, %v", c, err)
		}
		log.Println(status)
	}
}
