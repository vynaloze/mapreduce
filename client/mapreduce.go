package client

import (
	"github.com/vynaloze/mapreduce/api"
	"io"
	"log"
)

type Mapper interface {
	Map(key *api.Key, value *api.Value) <-chan *api.Pair
}

type Reducer interface {
	Reduce(key *api.Key, values <-chan *api.Value) <-chan *api.Pair
}

type MapperEntry struct {
	Name   string
	Mapper Mapper
}

type ReducerEntry struct {
	Name    string
	Reducer Reducer
}

func (c *client) RegisterMapper(entry MapperEntry) {
	c.mapper = &entry
	if c.reducer != nil {
		startMapReduceServer(c)
	}
}

func (c *client) RegisterReducer(entry ReducerEntry) {
	c.reducer = &entry
	if c.mapper != nil {
		startMapReduceServer(c)
	}
}

type mapReduceServer struct {
	api.UnimplementedMapReduceServer
	client *client
}

func (mr *mapReduceServer) Map(pair *api.Pair, stream api.MapReduce_MapServer) error {
	results := mr.client.mapper.Mapper.Map(pair.GetKey(), pair.GetValue())
	for pair := range results {
		if err := stream.Send(pair); err != nil {
			return err
		}
	}
	return nil
}

func (mr *mapReduceServer) Reduce(stream api.MapReduce_ReduceServer) error {
	pair, err := stream.Recv()

	values := make(chan *api.Value, 100)
	results := mr.client.reducer.Reducer.Reduce(pair.GetKey(), values)

	go func() {
		defer close(values)
		for {
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("error during internal reduce: %s", err)
			}
			values <- pair.GetValue()

			pair, err = stream.Recv()
		}
	}()
	for p := range results {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}
