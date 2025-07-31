# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/appointment-queue-service
RUN go mod download
RUN go build -o appointment-queue-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/appointment-queue-service/appointment-queue-service ./appointment-queue-service
COPY appointment-queue-service/.env .env
EXPOSE 8080 50053
CMD ["./appointment-queue-service"]
