package main

import (
	"flag"
	"github.com/vynaloze/mapreduce/engine/master"
	"github.com/vynaloze/mapreduce/engine/worker"
	"log"
)

func main() {
	masterAddr := flag.String("master", "", "Master node address. Empty means this node should be master")
	addr := flag.String("addr", ":50051", "Address to listen on")
	flag.Parse()

	if *masterAddr == "" {
		log.Println("starting as master on " + *addr)
		master.Run(*addr)
	} else {
		log.Println("starting as worker on " + *addr)
		worker.Run(*addr, *masterAddr)
	}
}
