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

// CandidateServiceServerImpl candidate service server type
type CandidateServiceServerImpl struct {
	Channel    *amqp.Channel
	Connection *mq.Connection
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

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerProfileQueue,
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

// SaveJobseekerProfilePhoto publish job seeker profile photo to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerProfilePhoto(context context.Context, in *domain.JobseekerProfilePhoto) (*domain.Error, error) {
	log.Println("SaveJobseekerProfilePhoto called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerProfilePhotoQueue,
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

// SaveJobseekerJobPreferences publish job seeker job preference to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerJobPreferences(context context.Context, in *domain.JobseekerJobPreferences) (*domain.Error, error) {
	log.Println("SaveJobseekerJobPreferences called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.UpdateJobseekerJobPreferencesQueue,
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

// SaveJobseekerSkills publish job seeker skills to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSkills(context context.Context, in *domain.JobseekerSkill) (*domain.Error, error) {
	log.Println("SaveJobseekerSkills called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerSkillsQueue,
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

// SaveJobseekerSummary publish job seeker summary to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerSummary(context context.Context, in *domain.JobseekerSummary) (*domain.Error, error) {
	log.Println("SaveJobseekerSummary called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.UpdateJobseekerSummaryQueue,
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

// SaveJobseekerEducation publish job seeker education to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	log.Println("SaveJobseekerEducation called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerEducationQueue,
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

// DeleteJobseekerEducation publish delete job seeker education to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerEducation(context context.Context, in *domain.JobseekerEducation) (*domain.Error, error) {
	log.Println("DeleteJobseekerEducation called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerEducationQueue,
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

// SaveJobseekerWorkExperience publish work experience to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	log.Println("SaveJobseekerWorkExperience called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerWorkExperienceQueue,
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

// DeleteJobseekerWorkExperience publish delete job seeker workexperience to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerWorkExperience(context context.Context, in *domain.JobseekerWorkExperience) (*domain.Error, error) {
	log.Println("DeleteJobseekerWorkExperience called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerWorkExperienceQueue,
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

// SaveJobseekerTraining publish job seeker training to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	log.Println("SaveJobseekerTraining called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerTrainingQueue,
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

// DeleteJobseekerTraining publish delete job seeker training to message queue
func (serviceImpl CandidateServiceServerImpl) DeleteJobseekerTraining(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	log.Println("DeleteJobseekerTraining called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerTrainingQueue,
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

// SaveJobseekerAffiliation publish job seeker affiliation to message queue
func (serviceImpl CandidateServiceServerImpl) SaveJobseekerAffiliation(context context.Context, in *domain.JobseekerTraining) (*domain.Error, error) {
	log.Println("SaveJobseekerAffiliation called from producer")

	b, _ := json.MarshalIndent(&in, "", "\t")
	m := mq.Message{
		Queue:         mq.SaveJobseekerAffiliationQueue,
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

func (CandidateServiceServerImpl) mustEmbedUnimplementedCandidateServiceServer() {}
