package domain

type JobQueueService interface {
	AddJob(job Job)
	GetJobs(status string) []Job
	ProcessJobs()
}
