package public

import (
	pb "github.com/vynaloze/mapreduce/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const JobLimit = 1

func New(lis net.Listener, s *grpc.Server) (<-chan *pb.Job, chan<- *pb.JobStatus) {
	return serve(lis, s)
}

type masterServer struct {
	pb.UnimplementedMasterServer
	jobs       chan<- *pb.Job
	jobsStatus <-chan *pb.JobStatus
}

func (s *masterServer) Submit(job *pb.Job, stream pb.Master_SubmitServer) error {
	s.jobs <- job
	for js := range s.jobsStatus {
		if err := stream.Send(js); err != nil {
			return err
		}
	}
	return nil
}

func serve(lis net.Listener, s *grpc.Server) (<-chan *pb.Job, chan<- *pb.JobStatus) {
	jobs := make(chan *pb.Job, JobLimit)
	jobsStatus := make(chan *pb.JobStatus)

	go func() {
		defer close(jobs)
		defer close(jobsStatus)

		pb.RegisterMasterServer(s, &masterServer{jobs: jobs, jobsStatus: jobsStatus})
		log.Printf("masterServer listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve masterServer: %v", err)
		}
	}()
	return jobs, jobsStatus
}
