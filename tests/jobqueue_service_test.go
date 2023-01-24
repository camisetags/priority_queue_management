package tests

import (
	"fmt"
	"priority_queue_management/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAndAddJobs(t *testing.T) {
	// Initialize JobQueueService
	jobQueueService := domain.NewJobQueueServiceImpl()

	// Test AddJob
	job := domain.Job{ID: 1, Data: []int{1, 2, 3}, Status: "pending"}
	jobQueueService.AddJob(job)
	assert.Equal(t, []domain.Job{job}, jobQueueService.GetJobs(""))

	// Test GetJobs
	assert.Equal(t, []domain.Job{job}, jobQueueService.GetJobs("pending"))
	assert.Equal(t, []domain.Job(nil), jobQueueService.GetJobs("processed"))
}

func TestProcessJobs(t *testing.T) {
	jobQueueService := domain.NewJobQueueServiceImpl()

	// Add jobs
	jobQueueService.AddJob(domain.Job{ID: 1, Data: []int{1, 2, 3}, Status: "pending"})
	jobQueueService.AddJob(domain.Job{ID: 2, Data: []int{4, 5, 6}, Status: "pending"})
	jobQueueService.AddJob(domain.Job{ID: 3, Data: []int{7, 8, 9}, Status: "pending"})
	jobQueueService.AddJob(domain.Job{ID: 4, Data: []int{1, 1, 1}, Status: "processed"})

	// Start processing jobs
	go jobQueueService.ProcessJobs()

	// Wait for jobs to be processed
	time.Sleep(time.Second * 10)

	// Check that the jobs have been processed
	for _, job := range jobQueueService.GetJobQueue().Jobs {
		if job.ID != 4 {
			sum := 0
			for _, number := range job.Data {
				sum += number
			}
			assert.Equal(t, "processed", job.Status)
			assert.Equal(t, fmt.Sprintf(`Test %d`, sum), job.Result, "should return the sum of the numbers")
		}
	}
}
