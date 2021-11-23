package client

import (
	"context"
	"github.com/vynaloze/mapreduce/api"
	"io"
	"log"
	"time"
)

func (c *client) SubmitAndWait(job *api.Job) {
	err := c.registerThisExecutable()
	if err != nil {
		log.Fatalf("cannot register mapreduce functions: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	stream, err := c.masterClient.Submit(ctx, job)
	if err != nil {
		log.Fatalf("could not submit: %v", err)
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
