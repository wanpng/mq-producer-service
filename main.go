package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
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

	router := mux.NewRouter()

	router.HandleFunc("/hc", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/", grpcServer.ServeHTTP)
	http.Handle("/", router)

	log.Printf("connect to http://localhost:%s/ for GRPC", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
