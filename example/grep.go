package main

import (
	"fmt"
	pb "github.com/vynaloze/mapreduce/api"
	"strings"
	"time"
)

type Grep struct {
	Pattern string
}

func (g *Grep) Map(key *pb.Key, value *pb.Value) <-chan *pb.Pair {
	// key: document name
	// value: document contents
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)

		for i, line := range strings.Split(value.GetValue(), "\n") {
			if strings.Contains(line, g.Pattern) {
				o <- &pb.Pair{Key: &pb.Key{Key: fmt.Sprintf("%s_%d", key.GetKey(), i)}, Value: &pb.Value{Value: line}}
			}
		}

		time.Sleep(1000 * time.Millisecond) // FIXME
	}()
	return o
}

type Identity struct{}

func (i *Identity) Reduce(key *pb.Key, values <-chan *pb.Value) <-chan *pb.Pair {
	// key: a word
	// values: a list of counts
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)
		time.Sleep(100 * time.Millisecond) //FIXME
		for value := range values {
			o <- &pb.Pair{Key: key, Value: value}
		}
	}()
	return o
}
