FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

WORKDIR /app/notification-service
RUN go mod download
RUN go build -o notification-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/notification-service/notification-service ./notification-service
COPY notification-service/.env .env

CMD ["./notification-service"]
