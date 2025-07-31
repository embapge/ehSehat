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
	ClinicDataClient   pb.ClinicDataServiceClient
	AppointmentClient  pb.AppointmentServiceClient
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

	clinicDataAddress := os.Getenv("CLINIC_DATA_SERVICE_URL")
	if clinicDataAddress == "" {
		clinicDataAddress = "localhost:50052"
	}
	if u, err := url.Parse(clinicDataAddress); err == nil && u.Host != "" {
		clinicDataAddress = u.Host
	}
	appointmentAddress := os.Getenv("APPOINTMENT_SERVICE_URL")
	if appointmentAddress == "" {
		appointmentAddress = "localhost:50053"
	}
	if u, err := url.Parse(appointmentAddress); err == nil && u.Host != "" {
		appointmentAddress = u.Host
	}

	authConn, err := grpc.DialContext(context.Background(), authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}

	consultationConn, err := grpc.DialContext(context.Background(), consultationAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Consultation Service: %v", err)
	}

	clinicDataConn, err := grpc.DialContext(context.Background(), clinicDataAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Clinic Data Service: %v", err)
	}

	appointmentConn, err := grpc.DialContext(context.Background(), appointmentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Appointment Service: %v", err)
	}

	return &GRPCClients{
		AuthClient:         pb.NewAuthServiceClient(authConn),
		ConsultationClient: pb.NewConsultationServiceClient(consultationConn),
		PaymentClient:      pb.NewPaymentServiceClient(consultationConn),
		ClinicDataClient:   pb.NewClinicDataServiceClient(clinicDataConn),
		AppointmentClient:  pb.NewAppointmentServiceClient(appointmentConn),
	}
}
