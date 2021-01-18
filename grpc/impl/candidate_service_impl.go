package impl

import (
	"context"
	"fmt"

	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"github.com/streadway/amqp"
	"github.com/wanpng/mq-producer-service/data/mq"
)

type CandidateServiceServerImpl struct {
	Channel *amqp.Channel
	service.CandidateServiceServer
}

func NewCandidateServiceServerImpl(ch *amqp.Channel) CandidateServiceServerImpl {
	return CandidateServiceServerImpl{
		Channel: ch,
	}
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerProfile(context context.Context, in *domain.JobseekerProfile) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerProfile called: %s %s\n", in.FirstName, in.LastName)

	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerProfile)

	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerJobPreferences(context context.Context, in *domain.JobseekerJobPreferences) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerJobPreferences called: %v\n", in)

	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.UpdateJobseekerJobPreferences)
	
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSkills(context context.Context, in *domain.JobseekerSkill) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerSkills called: %v\n", in)

	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerSkills)

	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSummary(context context.Context, in *domain.JobseekerSummary) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerSummary called: %v\n", in)

	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.UpdateJobseekerSummary)

	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerEducation called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	fmt.Printf("DeleteJobseekerEducation called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerWorkExperience called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	fmt.Printf("DeleteJobseekerWorkExperience called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	fmt.Printf("SaveJobseekerTraining called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	fmt.Printf("DeleteJobseekerTraining called: %v\n", in)
	return &domain.Error{
		Code: 0,
		Message: "",
	}, nil
}

func (CandidateServiceServerImpl) mustEmbedUnimplementedCandidateServiceServer() {}