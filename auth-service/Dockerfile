# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/auth-service
RUN go mod download
RUN go build -o auth-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/auth-service/auth-service ./auth-service
COPY auth-service/.env .env

EXPOSE 50051
CMD ["./auth-service"]
