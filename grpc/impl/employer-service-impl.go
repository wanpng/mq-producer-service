package impl

import (
	"context"

	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
)

// EmployerServiceServerImpl is a implementation of EmployerService
type EmployerServiceServerImpl struct {
	service.EmployerServiceServer
}

// SendToMQ function implementation of gRPC service
func (EmployerServiceServerImpl) sendEmployerMQ(context.Context, *domain.Employer) (*service.SendToMQResponse, error) {
	return nil, nil
}

// mustEmbedUnimplementedEmployerServiceServer function implementation of gRPC service
func (EmployerServiceServerImpl) mustEmbedUnimplementedEmployerServiceServer() {}