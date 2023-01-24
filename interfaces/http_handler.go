package interfaces

import (
	"net/http"
	"priority_queue_management/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	jobQueueService domain.JobQueueService
}

func setStatus(err error, c *fiber.Ctx) int {
	switch err.Error() {
	case "EOF":
		c.Status(http.StatusBadRequest)
		return http.StatusBadRequest
	case "Invalid request body":
		c.Status(http.StatusUnprocessableEntity)
		return http.StatusUnprocessableEntity
	case "Unprocessable Entity":
		c.Status(http.StatusUnprocessableEntity)
		return http.StatusUnprocessableEntity
	default:
		c.Status(http.StatusInternalServerError)
		return http.StatusInternalServerError
	}
}

func NewHTTPHandler(jobQueueService domain.JobQueueService) *HTTPHandler {
	return &HTTPHandler{
		jobQueueService: jobQueueService,
	}
}

func (h *HTTPHandler) GetJobs(c *fiber.Ctx) error {
	status := c.Query("status")
	jobs := h.jobQueueService.GetJobs(status)
	return c.JSON(jobs)
}

func (h *HTTPHandler) AddJob(c *fiber.Ctx) error {
	if c.Get("Authorization") != "allow" {
		return c.JSON(fiber.Map{"error": "Unauthorized", "status": http.StatusUnauthorized})
	}

	var job domain.Job
	if err := c.BodyParser(&job); err != nil {
		status := setStatus(err, c)
		return c.JSON(fiber.Map{"error": err.Error(), "status": status})
	}

	h.jobQueueService.AddJob(job)
	return c.JSON(job)
}
