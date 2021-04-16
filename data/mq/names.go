package mq

// QueueName type QueueName
type QueueName string

// ExchangeName type ExchangeName
type ExchangeName string

// RoutingKey type RoutingKey
type RoutingKey string

var (
	// Jobseeker jobseeker queue name
	Jobseeker QueueName = "jobseeker"
)

var (
	// JobseekerEx exchange name of jobseeker
	JobseekerEx ExchangeName = "jobseeker-ex"
	// JobsEx exchange name of jobs
	JobsEx ExchangeName = "jobs-ex"
)

var (
	// SaveJobseekerProfile save job seeker profile
	SaveJobseekerProfile RoutingKey = "save-jobseeker-profile"
	// SaveJobseekerProfilePhoto save job seeker profile photo
	SaveJobseekerProfilePhoto RoutingKey = "save-jobseeker-profile-photo"
	// UpdateJobseekerJobPreferences update job seeker job preferences
	UpdateJobseekerJobPreferences RoutingKey = "update-job-preferences"
	// SaveJobseekerSkills save job seeker skills
	SaveJobseekerSkills RoutingKey = "save-jobseeker-skills"
	// UpdateJobseekerSummary update jobseeker summary
	UpdateJobseekerSummary RoutingKey = "update-jobseeker-summary"
	// SaveJobseekerEducation save job seeker education
	SaveJobseekerEducation RoutingKey = "save-jobseeker-education"
	// DeleteJobseekerEducation delete job seeker education
	DeleteJobseekerEducation RoutingKey = "delete-jobseeker-education"
	// SaveJobseekerWorkExperience save job seeker work experience
	SaveJobseekerWorkExperience RoutingKey = "save-jobseeker-work-experience"
	// DeleteJobseekerWorkExperience delete job seeker work experience
	DeleteJobseekerWorkExperience RoutingKey = "delete-jobseeker-work-experience"
	// SaveJobseekerTraining save job seeker training
	SaveJobseekerTraining RoutingKey = "save-jobseeker-training"
	// DeleteJobseekerTraining delete job seeker training
	DeleteJobseekerTraining RoutingKey = "delete-jobseeker-training"

	// SaveJobInformation save job information
	SaveJobInformation RoutingKey = "save-job-information"
	// DeleteJoPost delete job post
	DeleteJoPost RoutingKey = "delete-job-post"
)

var (
	// JobseekerRoutingKeys job seeker routing keys
	JobseekerRoutingKeys = [...]RoutingKey{
		SaveJobseekerProfile,
		SaveJobseekerProfilePhoto,
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

	// JobsRoutingKeys jobs routing keys
	JobsRoutingKeys = [...]RoutingKey{
		SaveJobInformation,
		DeleteJoPost,
	}
)

var (
	SaveJobseekerProfileQueue          string = "save-jobseeker-profile-queue"
	SaveJobseekerProfilePhotoQueue     string = "save-jobseeker-profile-photo-queue"
	UpdateJobseekerJobPreferencesQueue string = "update-job-preferences-queue"
	SaveJobseekerSkillsQueue           string = "save-jobseeker-skills-queue"
	UpdateJobseekerSummaryQueue        string = "update-jobseeker-summary-queue"
	SaveJobseekerEducationQueue        string = "save-jobseeker-education-queue"
	DeleteJobseekerEducationQueue      string = "delete-jobseeker-education-queue"
	SaveJobseekerWorkExperienceQueue   string = "save-jobseeker-work-experience-queue"
	DeleteJobseekerWorkExperienceQueue string = "delete-jobseeker-work-experience-queue"
	SaveJobseekerTrainingQueue         string = "save-jobseeker-training-queue"
	DeleteJobseekerTrainingQueue       string = "delete-jobseeker-training-queue"

	SaveJobInformationQueue string = "save-job-information-queue"
	DeleteJoPostQueue       string = "delete-job-post-queue"
)
