package io

import (
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
)

type Split struct {
	Source string
	Offset int64
	Limit  int64
}

type Pair struct {
	Key   *internal.Key
	Value *internal.Value
}

type InputReader interface {
	Read(split Split) <-chan Pair
}

//type OutputWriter interface {
//	Write()
//}

type Handler interface {
	Split(spec *external.InputSpec) []Split
	InputReader
}
