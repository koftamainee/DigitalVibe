FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway main.go


FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]