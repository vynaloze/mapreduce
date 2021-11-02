package worker

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/vynaloze/mapreduce/engine/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	ttlSeconds       = 30
	backoffs         = 3
	heartbeatSeconds = ttlSeconds/backoffs - 1
)

func Run(addr, masterAddr string) {
	go startHeartbeat(addr, masterAddr)

	//TODO
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMapWorkerServer(s, &mapWorkerServer{addr: addr, taskPartitions: make(map[string]*taskPartitions)})
	log.Printf("mapWorkerServer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve mapWorkerServer: %v", err)
	}

	select {}
}

func startHeartbeat(addr, masterAddr string) {
	conn, err := grpc.Dial(masterAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRegistryClient(conn)

	clientId := uuid.New().String()
	fails := 0
	for {
		err := register(c, &pb.RegisterRequest{Addr: addr, Uuid: clientId, TtlSeconds: ttlSeconds})
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

func register(c pb.RegistryClient, req *pb.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.Register(ctx, req)
	return err
}
