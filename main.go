package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/soheilhy/cmux"
	"github.com/wanpng/mq-producer-service/config"
	"github.com/wanpng/mq-producer-service/data/mq"
	"github.com/wanpng/mq-producer-service/grpc/impl"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"google.golang.org/grpc"
)

func init() {
	config.InitSettings()
}

const defaultPort = "9000"

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

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

	mux := http.NewServeMux()
	th := http.HandlerFunc(healthCheckHandler)

	mux.Handle("/hc", th)

	httpServer := &http.Server{
		Handler: mux,
	}

	m := cmux.New(l)

	// grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.Any())

	go grpcServer.Serve(grpcL)
	go httpServer.Serve(httpL)

	log.Printf("Started grpc and http server at port: %s", port)

	m.Serve()
}
