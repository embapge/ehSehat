package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"ehSehat/auth-service/internal/auth/app"
	"ehSehat/auth-service/internal/auth/config"
	grpc2 "ehSehat/auth-service/internal/auth/delivery/grpc"
	authPb "ehSehat/auth-service/internal/auth/delivery/grpc/pb"
	"ehSehat/auth-service/internal/auth/infra"
	"ehSehat/auth-service/pkg/hasher"
	"ehSehat/auth-service/pkg/jwt"

	_ "ehSehat/auth-service/docs"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// @title Auth Service API
// @version 1.0
// @description Microservices Auth (DDD + Mongo + JWT) for Tokokecil
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := config.MySQLInit()
	userRepo := infra.NewMySQLUserRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	jwtManager := jwt.NewManager(jwtSecret)
	passwordHasher := hasher.NewBcrypt()
	authApp := app.NewAuthApp(userRepo, passwordHasher, jwtManager)

	grpcServer := grpc.NewServer()
	authGrpcHandler := grpc2.NewAuthHandler(authApp)
	authPb.RegisterAuthServiceServer(grpcServer, authGrpcHandler)

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Auth gRPC Service running at :" + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
