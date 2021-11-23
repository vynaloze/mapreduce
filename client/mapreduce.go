package client

import (
	"github.com/vynaloze/mapreduce/api"
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
	c.mappers = append(c.mappers, entry)
}

func (c *client) RegisterReducer(entry ReducerEntry) {
	c.reducers = append(c.reducers, entry)
}
