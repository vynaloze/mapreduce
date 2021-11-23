package worker

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	ttlSeconds       = 60
	backoffs         = 3
	heartbeatSeconds = 2
)

var id int
var mrs *mapReduceServer

func Run(addr, masterAddr string) {
	port, err := strconv.Atoi(strings.Split(addr, ":")[1])
	if err != nil {
		log.Fatalf("invalid addr: " + addr)
	}
	id = port % 100

	go startHeartbeat(addr, masterAddr)

	//TODO
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	internal.RegisterMapWorkerServer(s, &mapWorkerServer{addr: addr, taskPartitions: make(map[string]*taskPartitions)})
	log.Printf("mapWorkerServer listening at %v", lis.Addr())
	internal.RegisterReduceWorkerServer(s, &reduceWorkerServer{data: make(map[string][]string)})
	log.Printf("RegisterReduceWorkerServer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	select {}
}

func startHeartbeat(addr, masterAddr string) {
	conn, err := grpc.Dial(masterAddr, grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to master: %v", err)
	}
	defer conn.Close()
	c := internal.NewRegistryClient(conn)

	clientId := uuid.New().String()
	fails := 0
	for {
		rr := &internal.RegisterRequest{Addr: addr, Uuid: clientId, TtlSeconds: ttlSeconds}
		if mrs != nil {
			rr.Mappers = []string{mrs.mapper}
			rr.Reducers = []string{mrs.reducer}
		}
		err := register(c, rr)
		if err != nil {
			log.Printf("could not register: %v", err)
			fails++
			if fails >= backoffs {
				_ = conn.Close()
				log.Fatalf("max %d backoffs achieved - exiting", fails)
			}
		} else {
			fails = 0
		}
		time.Sleep(heartbeatSeconds * time.Second)
	}
}

func register(c internal.RegistryClient, req *internal.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	reply, err := c.Register(ctx, req)
	if err != nil {
		return err
	}
	if reply.Executable != nil {
		mrs = &mapReduceServer{mapper: reply.Mappers[0], reducer: reply.Reducers[0]}
		err = mrs.Persist(reply.Executable)
	}
	return err
}

type mapReduceServer struct {
	mapper  string
	reducer string

	client external.MapReduceClient
}

func (mr *mapReduceServer) Persist(e *external.MapReduceExecutable) error {
	name := fmt.Sprintf("/tmp/mapreduce_%d", id)
	log.Printf("saving mapreduce executable %s", name)
	err := os.WriteFile(name, e.Executable, 0744)
	if err != nil {
		return err
	}

	port := fmt.Sprintf("600%d", id)
	cmd := exec.Command(name, fmt.Sprintf("-mapreduce-port=%s", port))
	log.Printf("starting mapreduce executable %s", cmd.String())
	err = cmd.Start()
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second) // ehhh

	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to internal mapreduce: %v", err)
	}
	mr.client = external.NewMapReduceClient(conn)
	log.Printf("started internal mapreduce")
	return nil
}
