# Business Logic Service Dockerfile
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the application
RUN go build -o business-service .

# Final stage
FROM gcr.io/distroless/base

COPY --from=builder /app/business-service /business-service

# Expose the port the service will run on
EXPOSE 8082

# Command to run the executable
CMD ["/business-service"]
