package impl

import (
	"context"
	"log"

	"github.com/streadway/amqp"
	"github.com/wanpng/mq-producer-service/data/mq"
	"github.com/wanpng/mq-producer-service/grpc/domain"
	"github.com/wanpng/mq-producer-service/grpc/service"
)

// CandidateServiceServerImpl candidate service server type
type CandidateServiceServerImpl struct {
	Channel *amqp.Channel
	service.CandidateServiceServer
}

// NewCandidateServiceServerImpl create new instance of candidate service server
func NewCandidateServiceServerImpl(ch *amqp.Channel) CandidateServiceServerImpl {
	return CandidateServiceServerImpl{
		Channel: ch,
	}
}

// SaveJobseekerProfile publish job seeker profile to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerProfile(context context.Context, in *domain.JobseekerProfile) (*domain.Error, error) {
	log.Println("SaveJobseekerProfile called from producer")
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerProfile)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerJobPreferences publish job seeker job preference to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerJobPreferences(context context.Context, in *domain.JobseekerJobPreferences) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.UpdateJobseekerJobPreferences)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerSkills publish job seeker skills to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSkills(context context.Context, in *domain.JobseekerSkill) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerSkills)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerSummary publish job seeker summary to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSummary(context context.Context, in *domain.JobseekerSummary) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.UpdateJobseekerSummary)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerEducation publish job seeker education to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerEducation)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// DeleteJobseekerEducation publish delete job seeker education to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.DeleteJobseekerEducation)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerWorkExperience publish work experience to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerWorkExperience)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// DeleteJobseekerWorkExperience publish delete job seeker workexperience to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.DeleteJobseekerWorkExperience)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// SaveJobseekerTraining publish job seeker training to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.SaveJobseekerTraining)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

// DeleteJobseekerTraining publish delete job seeker training to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	mq.SendMQEx(in, serviceImpl.Channel, mq.JobseekerEx, mq.DeleteJobseekerTraining)

	return &domain.Error{
		Code:    0,
		Message: "",
	}, nil
}

func (CandidateServiceServerImpl) mustEmbedUnimplementedCandidateServiceServer() {}
