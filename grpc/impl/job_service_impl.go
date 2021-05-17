package impl

import (
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/wanpng/mq-producer-service/data/mq"
	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
)

// JobServiceServerImpl job service server implementation
type JobServiceServerImpl struct {
	Channel    *amqp.Channel
	Connection *mq.Connection
	service.JobServiceServer
}

// NewJobServiceServerImpl create new install of job service server
func NewJobServiceServerImpl(ch *amqp.Channel) JobServiceServerImpl {
	return JobServiceServerImpl{
		Channel: ch,
	}
}

// SendJobToMQ publish job to message queue
func (serviceImpl JobServiceServerImpl) SendJobToMQ(context context.Context, in *domain.Job) (*domain.Error, error) {
	log.Println("SendJobToMQ called from producer")
	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobInformationQueue,
		ReplyTo:       "",
		ContentType:   "application/json",
		CorrelationID: "",
		Priority:      1,
		Body:          mq.MessageBody{Data: b, Type: ""},
	}

	if err := serviceImpl.Connection.Publish(m); err != nil {
		log.Fatalf("Publishing error %s", err)
	}

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// DeleteJobPost delete job to message queue
func (serviceImpl JobServiceServerImpl) DeleteJobPost(context context.Context, in *domain.DeleteJob) (*domain.Error, error) {
	log.Println("DeleteJobPost called from producer")
	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.DeleteJobPostQueue,
		ReplyTo:       "",
		ContentType:   "application/json",
		CorrelationID: "",
		Priority:      1,
		Body:          mq.MessageBody{Data: b, Type: ""},
	}

	if err := serviceImpl.Connection.Publish(m); err != nil {
		log.Fatalf("Publishing error %s", err)
	}

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// UpdateJobCompanyProfile delete job to message queue
func (serviceImpl JobServiceServerImpl) UpdateJobCompanyProfile(context context.Context, in *domain.JobCompanyProfile) (*domain.Error, error) {
	log.Println("UpdateJobCompanyProfile called from producer")
	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.UpdateJobCompanyProfileQueue,
		ReplyTo:       "",
		ContentType:   "application/json",
		CorrelationID: "",
		Priority:      1,
		Body:          mq.MessageBody{Data: b, Type: ""},
	}

	if err := serviceImpl.Connection.Publish(m); err != nil {
		log.Fatalf("Publishing error %s", err)
	}

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

func (JobServiceServerImpl) mustEmbedUnimplementedJobServiceServer() {}
