package client

import (
	"github.com/vynaloze/mapreduce/api"
	"google.golang.org/grpc"
	"log"
)

type Client interface {
	RegisterMapper(entry MapperEntry)
	RegisterReducer(entry ReducerEntry)

	SubmitAndWait(job *api.Job)
}

type client struct {
	mapper  *MapperEntry
	reducer *ReducerEntry

	masterClient            api.MasterClient
	mapReduceRegistryClient api.MapReduceRegistryClient
}

func New() Client {
	conn, err := grpc.Dial(":50050", grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	//defer conn.Close()

	return &client{
		masterClient:            api.NewMasterClient(conn),
		mapReduceRegistryClient: api.NewMapReduceRegistryClient(conn),
	}
}
