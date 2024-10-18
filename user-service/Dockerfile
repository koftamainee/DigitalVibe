# User Management Service Dockerfile
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the application
RUN go build -o user-service .

# Final stage
FROM gcr.io/distroless/base

COPY --from=builder /app/user-service /user-service

# Expose the port the service will run on
EXPOSE 8081

# Command to run the executable
CMD ["/user-service"]
