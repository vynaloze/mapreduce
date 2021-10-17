package main

import "flag"

func main() {
	masterAddr := flag.String("master", "", "Master node address. Empty means this node should be master")
	addr := flag.String("addr", ":50051", "Address to listen on")
	flag.Parse()

	if *masterAddr == "" {
		serveMaster(*addr)
	} else {
		serveWorker(*masterAddr)
	}
}
