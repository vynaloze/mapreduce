package main

import (
	"flag"
	"github.com/vynaloze/mapreduce/engine/master"
	"github.com/vynaloze/mapreduce/engine/worker"
)

func main() {
	masterAddr := flag.String("master", "", "Master node address. Empty means this node should be master")
	addr := flag.String("addr", ":50051", "Address to listen on")
	flag.Parse()

	if *masterAddr == "" {
		master.Run(*addr)
	} else {
		worker.Run(*addr, *masterAddr)
	}
}
