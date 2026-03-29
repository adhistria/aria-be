
# Aria-BE

Proof of Concept for Aria Backend Service.

## Prerequisites

- Go 1.23 or higher
- Docker (for containerized setup)
- GCP CLI (for deployment)

## Running Locally with Go

```bash
# Clone the repository
git clone <repository-url>
cd aria-be

# Install dependencies
go mod download

# Run the application
go run main.go
```

The service will start on `http://localhost:8080`

## Running with Docker

```bash
# Build the Docker image
docker build -t aria-be .

# Run the container
docker run -p 8080:8080 aria-be
```

## Accessing Swagger Documentation

Once the service is running, access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

For GCP Cloud Run deployments:

```
https://YOUR_CLOUD_RUN_URL/swagger/index.html
```
