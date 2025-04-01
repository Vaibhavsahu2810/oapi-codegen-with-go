# License API

A RESTful API in Go for managing software licenses, generated from an OpenAPI 3.0 specification using [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen). Built with the Gin web framework and includes Swagger UI for API documentation.

## Features

- OpenAPI 3.0 spec-driven development
- Auto-generated type-safe server handlers using `oapi-codegen`
- Gin-based HTTP server
- Create/Update License endpoints
- Integrated Swagger UI for interactive documentation



##  Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Vaibhavsahu2810/oapi-codegen-with-go.git
cd oapi-codegen-with-go
```

### Install dependencies
go mod tidy

### Ensure oapi-codegen is installed:

go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

## Run the server

go run main.go

### Access the API docs
Visit: http://localhost:8080/swagger-ui/
