# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/gateway-service
RUN go mod download
RUN go build -o gateway-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/gateway-service/gateway-service ./gateway-service
COPY gateway-service/.env .env
COPY gateway-service/docs ./docs

EXPOSE 8000
CMD ["./gateway-service"]
