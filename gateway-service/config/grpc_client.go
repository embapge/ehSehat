package config

import (
	"context"
	pb "ehSehat/gateway-service/handler/pb" // adjust if consultation proto is in different pb
	"log"
	"net/url"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	AuthClient         pb.AuthServiceClient
	ConsultationClient pb.ConsultationServiceClient
	PaymentClient      pb.PaymentServiceClient
}

func NewGRPCClients() *GRPCClients {
	authAddress := os.Getenv("AUTH_SERVICE_URL")
	if authAddress == "" {
		authAddress = "localhost:50051"
	}
	if u, err := url.Parse(authAddress); err == nil && u.Host != "" {
		authAddress = u.Host
	}

	consultationAddress := os.Getenv("CONSULTATION_SERVICE_URL")
	if consultationAddress == "" {
		consultationAddress = "localhost:50054"
	}
	if u, err := url.Parse(consultationAddress); err == nil && u.Host != "" {
		consultationAddress = u.Host
	}

	authConn, err := grpc.DialContext(context.Background(), authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}

	consultationConn, err := grpc.DialContext(context.Background(), consultationAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Consultation Service: %v", err)
	}

	return &GRPCClients{
		AuthClient:         pb.NewAuthServiceClient(authConn),
		ConsultationClient: pb.NewConsultationServiceClient(consultationConn),
		PaymentClient:      pb.NewPaymentServiceClient(consultationConn),
	}
}
