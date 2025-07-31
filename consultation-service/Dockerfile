# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/consultation-service
RUN go mod download
RUN go build -o consultation-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/consultation-service/consultation-service ./consultation-service
# COPY consultation-service/.env .env

EXPOSE 50054
CMD ["./consultation-service"]
