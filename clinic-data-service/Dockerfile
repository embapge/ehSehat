# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/clinic-data-service
RUN go mod download
RUN go build -o clinic-data-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/clinic-data-service/clinic-data-service ./clinic-data-service
COPY clinic-data-service/.env .env

EXPOSE 50052
CMD ["/app/clinic-data-service"]
