package main

import (
	"priority_queue_management/domain"
	"priority_queue_management/interfaces"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	jobQueueService := domain.NewJobQueueServiceImpl()
	go jobQueueService.ProcessJobs()

	httpHandler := interfaces.NewHTTPHandler(jobQueueService)

	app := fiber.New()

	app.Get("/jobs", httpHandler.GetJobs)
	app.Post("/jobs", httpHandler.AddJob)

	app.Listen(":3000")
}
