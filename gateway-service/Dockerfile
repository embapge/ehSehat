FROM golang:1.24-alpine AS build

# Copy root module files and proto for replace directive resolution
WORKDIR /app
COPY go.mod go.sum ./
COPY proto ./proto

# Prepare gateway-service module context
WORKDIR /app/gateway-service
# Copy module files
COPY gateway-service/go.mod gateway-service/go.sum ./
# Copy service source
COPY gateway-service/. ./

# Download dependencies and build
RUN go mod download && \
    go build -o gateway-service cmd/main.go

FROM alpine:latest
WORKDIR /app
# Copy only the built binary with .bin suffix to avoid directory name conflict
COPY --from=build /app/gateway-service/gateway-service ./gateway-service.bin
# Copy service env file
COPY gateway-service/.env .env

EXPOSE 8000
CMD ["./gateway-service.bin"]