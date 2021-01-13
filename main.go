package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"github.com/wanpng/mq-producer-service/grpc/impl"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	employerServiceImpl := impl.EmployerServiceServerImpl{}

	grpcServer := grpc.NewServer()

	service.RegisterEmployerServiceServer(grpcServer, employerServiceImpl)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}