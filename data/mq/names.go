package mq

type QueueName string
type ExchangeName string
type RoutingKey string

var (
	Jobseeker QueueName = "jobseeker"
)

var (
	JobseekerEx ExchangeName = "jobseeker-ex"
)

var (
	SaveJobseekerProfile RoutingKey = "save-jobseeker-profile"
	UpdateJobseekerJobPreferences RoutingKey = "update-job-preferences"
	SaveJobseekerSkills RoutingKey = "save-jobseeker-skills"
	UpdateJobseekerSummary RoutingKey = "update-jobseeker-summary"
	SaveJobseekerEducation RoutingKey = "save-jobseeker-education"
	DeleteJobseekerEducation RoutingKey = "delete-jobseeker-education"
	SaveJobseekerWorkExperience RoutingKey = "save-jobseeker-work-experience"
	DeleteJobseekerWorkExperience RoutingKey = "delete-jobseeker-work-experience"
	SaveJobseekerTraining RoutingKey = "save-jobseeker-training"
	DeleteJobseekerTraining RoutingKey = "delete-jobseeker-training"
)

var (
	JobseekerRoutingKeys = [...]RoutingKey{
		SaveJobseekerProfile,
		UpdateJobseekerJobPreferences,
		SaveJobseekerSkills,
		UpdateJobseekerSummary,
		SaveJobseekerEducation,
		DeleteJobseekerEducation,
		SaveJobseekerWorkExperience,
		DeleteJobseekerWorkExperience,
		SaveJobseekerTraining,
		DeleteJobseekerTraining,
	}
)