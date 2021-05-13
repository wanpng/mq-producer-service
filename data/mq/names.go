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
	SaveJobseekerProfile RoutingKey = "save-jobseeker-profile-1"
	// SaveJobseekerProfilePhoto save job seeker profile photo
	SaveJobseekerProfilePhoto RoutingKey = "save-jobseeker-profile-photo-1"
	// UpdateJobseekerJobPreferences update job seeker job preferences
	UpdateJobseekerJobPreferences RoutingKey = "update-job-preferences-1"
	// SaveJobseekerSkills save job seeker skills
	SaveJobseekerSkills RoutingKey = "save-jobseeker-skills-1"
	// UpdateJobseekerSummary update jobseeker summary
	UpdateJobseekerSummary RoutingKey = "update-jobseeker-summary-1"
	// SaveJobseekerEducation save job seeker education
	SaveJobseekerEducation RoutingKey = "save-jobseeker-education-1"
	// DeleteJobseekerEducation delete job seeker education
	DeleteJobseekerEducation RoutingKey = "delete-jobseeker-education-1"
	// SaveJobseekerWorkExperience save job seeker work experience
	SaveJobseekerWorkExperience RoutingKey = "save-jobseeker-work-experience-1"
	// DeleteJobseekerWorkExperience delete job seeker work experience
	DeleteJobseekerWorkExperience RoutingKey = "delete-jobseeker-work-experience-1"
	// SaveJobseekerTraining save job seeker training
	SaveJobseekerTraining RoutingKey = "save-jobseeker-training-1"
	// DeleteJobseekerTraining delete job seeker training
	DeleteJobseekerTraining RoutingKey = "delete-jobseeker-training-1"

	// SaveJobInformation save job information
	SaveJobInformation RoutingKey = "save-job-information-1"
	// DeleteJobPost delete job post
	DeleteJobPost RoutingKey = "delete-job-post-1"
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
		DeleteJobPost,
	}
)

var (
	SaveJobseekerProfileQueue          string = "save-jobseeker-profile-queue-1"
	SaveJobseekerProfilePhotoQueue     string = "save-jobseeker-profile-photo-queue-1"
	UpdateJobseekerJobPreferencesQueue string = "update-job-preferences-queue-1"
	SaveJobseekerSkillsQueue           string = "save-jobseeker-skills-queue-1"
	UpdateJobseekerSummaryQueue        string = "update-jobseeker-summary-queue-1"
	SaveJobseekerEducationQueue        string = "save-jobseeker-education-queue-1"
	DeleteJobseekerEducationQueue      string = "delete-jobseeker-education-queue-1"
	SaveJobseekerWorkExperienceQueue   string = "save-jobseeker-work-experience-queue-1"
	DeleteJobseekerWorkExperienceQueue string = "delete-jobseeker-work-experience-queue-1"
	SaveJobseekerTrainingQueue         string = "save-jobseeker-training-queue-1"
	DeleteJobseekerTrainingQueue       string = "delete-jobseeker-training-queue-1"

	SaveJobInformationQueue string = "save-job-information-queue-1"
	DeleteJobPostQueue      string = "delete-job-post-queue-1"
)
