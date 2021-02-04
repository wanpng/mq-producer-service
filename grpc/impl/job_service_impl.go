package impl

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/wanpng/mq-producer-service/data/mq"
	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
)

// JobServiceServerImpl job service server implementation
type JobServiceServerImpl struct {
	Channel *amqp.Channel
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
	fmt.Printf("SendJobToMQ called")
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobsEx, mq.SaveJobInformation)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

func (JobServiceServerImpl) mustEmbedUnimplementedJobServiceServer() {}
