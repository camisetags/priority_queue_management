# Priority Queue Management

This project is a simple priority queue management system that allows users to add and retrieve jobs from a priority queue, using a REST API.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites
- Go version 1.14 or later
- Fiber v2
- Testify

## Installing
Clone the repository and navigate to the project directory

```bash
git clone https://github.com/camisetags/priority-queue-management.git
cd priority-queue-management
```
## Running the application
To start the application, run the following command from the project directory:

```bash
go run main.go
```

By default, the application will listen on port 3000.
You can change the port by setting an environment variable PORT, before running the application:

```bash
export PORT=8080
```

## Testing
To run the tests, navigate to the project directory and run the following command:

```bash
go test -v ./...
```

## Endpoints
The following endpoints are available:

### GET /jobs
Retrieves all jobs in the queue.
You can filter by status by passing the status parameter.

### POST /jobs
Adds a job to the queue.
The priority of the queue will be "first in first out" (FIFO), the ID will be the parameter representing the priority.

## Built With
- Fiber - The web framework used
- Testify - The testing toolkit

## Authors
Gabriel Seixas
