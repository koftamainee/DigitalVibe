# Stage 1: Build the Go app
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o business-service main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/business-service .

EXPOSE 8082

CMD ["./business-service"]