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

	if err != nil {
		log.Fatalf("Listen error in %s", err)
	}

	conn := mq.NewConnection("my-producer", "my-exchange",
		mq.SaveJobseekerProfileQueue,
		mq.SaveJobseekerProfilePhotoQueue,
		mq.UpdateJobseekerJobPreferencesQueue,
		mq.SaveJobseekerSkillsQueue,
		mq.UpdateJobseekerSummaryQueue,
		mq.SaveJobseekerEducationQueue,
		mq.DeleteJobseekerEducationQueue,
		mq.SaveJobseekerWorkExperienceQueue,
		mq.DeleteJobseekerWorkExperienceQueue,
		mq.SaveJobseekerTrainingQueue,
		mq.DeleteJobseekerTrainingQueue,
		mq.SaveJobseekerAffiliationQueue,
		mq.SaveJobInformationQueue,
		mq.DeleteJobPostQueue,
		mq.UpdateJobCompanyProfileQueue,
	)

	if err := conn.Connect(); err != nil {
		panic(err)
	}

	if err := conn.BindQueue(); err != nil {
		panic(err)
	}

	jobServiceImpl := impl.JobServiceServerImpl{
		Connection: conn,
	}
	candidateServiceImpl := impl.CandidateServiceServerImpl{
		Connection: conn,
	}

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

	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.Any())

	go grpcServer.Serve(grpcL)
	go httpServer.Serve(httpL)

	log.Printf("Started grpc and http server at port: %s", port)

	m.Serve()
}
