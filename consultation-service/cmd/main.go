package main

import (
	"ehSehat/consultation-service/config"
	"ehSehat/consultation-service/internal/consultation/app"
	grpc2 "ehSehat/consultation-service/internal/consultation/delivery/grpc"
	"ehSehat/consultation-service/internal/consultation/infra"
	consultationPb "ehSehat/proto/consultation"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatal("gRPC port is empty")
	}

	db := config.MongoDB()

	consultationCollection := config.GetCollection(db, "consultations")
	consultationRepo := infra.NewMongoConsultation(consultationCollection)
	consultationApp := app.NewConsultationApp(consultationRepo)
	consultationHandler := grpc2.NewConsultationHandler(consultationApp)

	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen gRPC: %v", err)
		}

		s := grpc.NewServer()

		consultationPb.RegisterConsultationServiceServer(s, consultationHandler)
		log.Println("gRPC Consultation Service running at:" + grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Block main goroutine - ga langsung exit
	select {}
}
