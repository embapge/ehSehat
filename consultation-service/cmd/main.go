package main

import (
	"ehSehat/consultation-service/config"
	"ehSehat/consultation-service/internal/consultation/app"
	grpc2 "ehSehat/consultation-service/internal/consultation/delivery/grpc"
	consultationPb "ehSehat/consultation-service/internal/consultation/delivery/grpc/pb"
	"ehSehat/consultation-service/internal/consultation/domain"
	"ehSehat/consultation-service/internal/consultation/infra"
	paymentApp "ehSehat/consultation-service/internal/payment/app"
	grpc3 "ehSehat/consultation-service/internal/payment/delivery/grpc"
	paymentPb "ehSehat/consultation-service/internal/payment/delivery/grpc/pb"
	paymentInfra "ehSehat/consultation-service/internal/payment/infra"
	"ehSehat/libs/utils/rabbitmqown"
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

	mongo := config.MongoDB()
	mySQL := config.MySQLDB()

	conn, ch, err := rabbitmqown.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}()

	queueName := os.Getenv("RABBIT_QUEUE")
	_, err = rabbitmqown.DeclareQueue(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	consultationCollection := config.GetCollection(mongo, "consultations")
	consultationRepo := infra.NewMongoConsultation(consultationCollection)
	consultationApp := app.NewConsultationApp(consultationRepo)
	consultationHandler := grpc2.NewConsultationHandler(consultationApp, ch)

	paymentRepo := paymentInfra.NewMySQLPayment(mySQL)
	gateway := paymentInfra.NewXendit()
	var consultationService domain.ConsultationService = consultationApp
	appPayment := paymentApp.NewPaymentApp(paymentRepo, consultationService, gateway)
	paymentHandler := grpc3.NewPaymentHandler(appPayment, ch)

	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen gRPC: %v", err)
		}

		s := grpc.NewServer()

		paymentPb.RegisterPaymentServiceServer(s, paymentHandler)
		consultationPb.RegisterConsultationServiceServer(s, consultationHandler)
		log.Println("gRPC Consultation Service running at:" + grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	select {}
}
