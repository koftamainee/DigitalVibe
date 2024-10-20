# Stage 1: Build the Go app
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o task-service main.go


FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/task-service .

EXPOSE 8082

CMD ["./task-service"]