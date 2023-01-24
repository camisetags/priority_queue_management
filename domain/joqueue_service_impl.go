package domain

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type JobQueueServiceImpl struct {
	jobQueue *JobQueue
	mu       sync.RWMutex
}

func NewJobQueueServiceImpl() *JobQueueServiceImpl {
	return &JobQueueServiceImpl{
		jobQueue: &JobQueue{},
	}
}

func (jqs *JobQueueServiceImpl) GetJobQueue() *JobQueue {
	return jqs.jobQueue
}

func (jqs *JobQueueServiceImpl) AddJob(job Job) {
	jqs.mu.Lock()
	defer jqs.mu.Unlock()
	jqs.jobQueue.Jobs = append(jqs.jobQueue.Jobs, job)
}

func (jqs *JobQueueServiceImpl) GetJobs(status string) []Job {
	jqs.mu.RLock()
	defer jqs.mu.RUnlock()
	var jobs []Job
	for _, job := range jqs.jobQueue.Jobs {
		if job.Status == status || status == "" {
			jobs = append(jobs, job)
		}
	}
	return jobs
}

func (jqs *JobQueueServiceImpl) ProcessJobs() {
	for {
		jqs.mu.Lock()
		for i, job := range jqs.jobQueue.Jobs {
			if job.Status != "processed" {
				result := 0
				for _, value := range job.Data {
					result += value
				}
				fmt.Printf("Job ID: %d, Result: %d\n", job.ID, result)
				jqs.jobQueue.Jobs[i].Status = "processed"
				jqs.jobQueue.Jobs[i].Result = fmt.Sprintf("Test %d", result)
			}
		}
		jqs.mu.Unlock()
		time.Sleep(time.Duration(getSeconds()) * time.Second)
	}
}

func getSeconds() int {
	interval := 10
	envInterval := os.Getenv("JOB_PROCESSING_INTERVAL")
	if envInterval != "" {
		intVal, err := strconv.Atoi(envInterval)
		if err == nil {
			interval = intVal
		}
	}
	return interval
}
