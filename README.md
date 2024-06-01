# go-rabbitmq-docker

This repository showcases a project using Go (Golang), RabbitMQ, and Docker, built with a layered architecture. It includes a Fibonacci sequence demonstration to illustrate the interaction between an API and a consumer. The API, built using the Gin framework, accepts a number as input, and the consumer processes the calculation.

## Getting Started

To get this project up and running, follow these steps:

### 1. Start the Project

Run the following command to start the Docker containers:

```sh
docker-compose up -d
```

### 2. Publish an Event

Use the following `curl` request to publish an event:

```sh
curl --location 'http://127.0.0.1:8000/api/publish' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'number=1000'
```

### 3. Run the Consumer

To run the consumer, first, get shell access to the Go container and then execute the consumer command:

```sh
# Access the Go container
docker exec -it rabbitmq_go sh

# Run the consumer command
go run cli/command.go fibonacci
```

### 4. Run Tests

To run tests, use the following command inside the `test` directory:

```sh
go test ./... -v
```

---

**Keywords:** go, docker, rabbitmq, gin

---

By following the steps above, you can effectively use the functionality provided by this project.