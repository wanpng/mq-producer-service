package impl

import (
	"context"

	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"github.com/streadway/amqp"
	"github.com/wanpng/mq-producer-service/data/mq"
)

// EmployerServiceServerImpl is a implementation of EmployerService
type EmployerServiceServerImpl struct {
	Channel *amqp.Channel
	service.EmployerServiceServer
}

func NewEmployerServiceServerImpl(ch *amqp.Channel) EmployerServiceServerImpl {
	return EmployerServiceServerImpl{
		Channel: ch,
	}
}

// SendToMQ function implementation of gRPC service
func (serviceImpl EmployerServiceServerImpl) SendEmployerMQ(context context.Context, in *domain.Employer) (*service.SendToMQResponse, error) {
	mq.SendMQ(in, serviceImpl.Channel, mq.Employer)

	println("mq sent")
	return &service.SendToMQResponse{
		Employer: in,
		Error: nil,
	}, nil
}

// mustEmbedUnimplementedEmployerServiceServer function implementation of gRPC service
func (EmployerServiceServerImpl) mustEmbedUnimplementedEmployerServiceServer() {}