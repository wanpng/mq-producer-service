package main

import (
	"log"
	"net"

	"github.com/wanpng/mq-producer-service/config"
	"github.com/wanpng/mq-producer-service/data/mq"
	"github.com/wanpng/mq-producer-service/grpc/impl"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"google.golang.org/grpc"
)

func init() {
	config.InitSettings()
}

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	conn, ch, err := mq.DeclareExchange(mq.JobseekerEx)

	if err != nil {
		log.Fatalf("Unable to declare queue")
	}

	defer conn.Close()
	defer ch.Close()

	jobServiceImpl := impl.NewJobServiceServerImpl(ch)
	candidateServiceImpl := impl.NewCandidateServiceServerImpl(ch)

	grpcServer := grpc.NewServer()

	service.RegisterJobServiceServer(grpcServer, jobServiceImpl)
	service.RegisterCandidateServiceServer(grpcServer, candidateServiceImpl)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
