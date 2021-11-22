package main

import (
	pb "github.com/vynaloze/mapreduce/api"
	"strconv"
	"strings"
	"time"
)

type WordCount struct {}

func (w *WordCount) Map(key *pb.Key, value *pb.Value) <-chan *pb.Pair {
	// key: document name
	// value: document contents
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)

		for _, word := range strings.Fields(value.GetValue()) {
			k := strings.ReplaceAll(word, ",", "")
			kk := strings.ReplaceAll(k, ".", "")

			o <- &pb.Pair{Key: &pb.Key{Key: kk}, Value: &pb.Value{Value: "1"}}
		}
		time.Sleep(2000 * time.Millisecond) // FIXME
	}()
	return o
}


type Adder struct {}

func (a *Adder) Reduce(key *pb.Key, values <-chan *pb.Value) <-chan *pb.Pair {
	// key: a word
	// values: a list of counts
	o := make(chan *pb.Pair)
	go func() {
		defer close(o)

		var result int
		for v := range values {
			i, err := strconv.Atoi(v.GetValue())
			if err != nil {
				panic(err)
			}
			result += i
		}
		time.Sleep(100 * time.Millisecond) //FIXME

		o <- &pb.Pair{Key: key, Value: &pb.Value{Value: strconv.Itoa(result)}}
	}()
	return o
}
