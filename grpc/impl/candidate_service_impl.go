package impl

import (
	"context"

	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
	"github.com/streadway/amqp"
	// "github.com/wanpng/mq-producer-service/data/mq"
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

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerProfile(context.Context, *domain.JobseekerProfile) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerJobPreferences(context.Context, *domain.JobseekerJobPreferences) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSkills(context.Context, *domain.JobseekerSkill) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSummary(context.Context, *domain.JobseekerSummary) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerEducation(context.Context, *domain.JobseekerEducation) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerEducation(context.Context, *domain.JobSeekerEducationDelete) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerWorkExperience(context.Context, *domain.JobseekerWorkExperience) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerWorkExperience(context.Context, *domain.JobseekerWorkExperienceDelete) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) SaveJobseekerTraining(context.Context, *domain.JobseekerTraining) (*domain.Error, error) {
	return nil, nil
}

func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerTraining(context.Context, *domain.JobseekerTrainingDelete) (*domain.Error, error) {
	return nil, nil
}

func (CandidateServiceServerImpl) mustEmbedUnimplementedCandidateServiceServer() {}