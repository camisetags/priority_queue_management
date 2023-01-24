package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"priority_queue_management/domain"
	"priority_queue_management/interfaces"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestEmptyJobs(t *testing.T) {
	jobQueueService := domain.NewJobQueueServiceImpl()

	httpHandler := interfaces.NewHTTPHandler(jobQueueService)

	app := fiber.New()

	app.Get("/jobs", httpHandler.GetJobs)
	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "null", bodyString)
}

func TestAddJobs(t *testing.T) {
	jobQueueService := domain.NewJobQueueServiceImpl()

	httpHandler := interfaces.NewHTTPHandler(jobQueueService)

	app := fiber.New()
	job := domain.Job{ID: 1, Data: []int{1, 2, 3}, Status: "pending"}

	app.Post("/jobs", httpHandler.AddJob)

	jobJSON, _ := json.Marshal(job)

	req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(jobJSON))
	req.Header.Add("Authorization", "allow")
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `{"id":1,"data":[1,2,3],"status":"pending"}`, bodyString)
}

func TestGetJobs(t *testing.T) {
	jobQueueService := domain.NewJobQueueServiceImpl()

	httpHandler := interfaces.NewHTTPHandler(jobQueueService)

	app := fiber.New()

	job := domain.Job{ID: 1, Data: []int{1, 2, 3}, Status: "pending"}
	jobQueueService.AddJob(job)

	app.Get("/jobs", httpHandler.GetJobs)

	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `[{"id":1,"data":[1,2,3],"status":"pending"}]`, bodyString)
}
