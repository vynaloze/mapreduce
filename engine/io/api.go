package io

import (
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
)

type InputReader interface {
	Read(split *internal.Split) <-chan *external.Pair
}

type OutputWriter interface {
	Write(pairs <-chan *external.Pair)
}

type Handler interface {
	Split(spec *external.InputSpec) []internal.Split
	InputReader
	OutputWriter
}
