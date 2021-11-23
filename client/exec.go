package client

import (
	"context"
	"flag"
	"github.com/vynaloze/mapreduce/api"
	"log"
	"os"
	"time"
)

func (c *client) registerThisExecutable() error {
	mapperNames, reducerNames := make([]string, 0), make([]string, 0)
	for _, m := range c.mappers {
		mapperNames = append(mapperNames, m.Name)
	}
	for _, r := range c.reducers {
		reducerNames = append(reducerNames, r.Name)
	}
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	executableBytes, err := os.ReadFile(executable)

	msg := api.MapReduceExecutable{
		Mappers:    mapperNames,
		Reducers:   reducerNames,
		Executable: executableBytes,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err = c.mapReduceRegistryClient.Register(ctx, &msg)
	return err
}

func init() {
	var port int
	flag.IntVar(&port, "mapreduce-port", 0, "port to serve mapreduce on")
	flag.Parse()

	if port != 0 {
		log.Printf("starting mapreduce on %d", port)
		//TODO start mapreduce
	}
}
