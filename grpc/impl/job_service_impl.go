package impl

import (
	"context"

	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"github.com/streadway/amqp"
	// "github.com/wanpng/mq-producer-service/data/mq"
)

type JobServiceServerImpl struct {
	Channel *amqp.Channel
	service.JobServiceServer
}

func NewJobServiceServerImpl(ch *amqp.Channel) JobServiceServerImpl {
	return JobServiceServerImpl{
		Channel: ch,
	}
}

func (serviceImpl JobServiceServerImpl) SendJobToMQ(context context.Context, in *domain.Job) (*service.SendJobToMQResponse, error) {
	return nil, nil
}

func (JobServiceServerImpl) mustEmbedUnimplementedJobServiceServer() {}