package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"appointment-queue-service/config"

	// Appointment
	appointmentApp "appointment-queue-service/internal/appointment/app"
	appointmentGRPC "appointment-queue-service/internal/appointment/delivery/grpc"
	appointmentHTTP "appointment-queue-service/internal/appointment/delivery/http"
	appointmentRepo "appointment-queue-service/internal/appointment/domain"

	// Queue
	queueApp "appointment-queue-service/internal/queue/app"
	queueHTTP "appointment-queue-service/internal/queue/delivery/http"
	queueRepo "appointment-queue-service/internal/queue/domain"

	"appointment-queue-service/internal/appointment/delivery/grpc/pb"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load .env or OS env
	config.LoadEnv()

	// DB
	db := config.DBInit()

	// --- Queue ---
	qRepo := queueRepo.NewQueueRepository(db)
	qService := queueRepo.NewQueueService(qRepo)
	qApp := queueApp.NewQueueApp(qService)

	// --- Appointment ---
	aRepo := appointmentRepo.NewAppointmentRepository(db)
	aService := appointmentRepo.NewAppointmentService(aRepo)
	aApp := appointmentApp.NewAppointmentApp(aService)

	// --- HTTP Router ---
	router := httprouter.New()
	queueHTTP.NewQueueHandler(router, qApp)
	appointmentHTTP.NewAppointmentHandler(router, aApp)

	// Get ports from env
	httpPort := config.GetEnv("HTTP_PORT", "8080")
	grpcPort := config.GetEnv("GRPC_PORT", "50051")

	// --- gRPC Server ---
	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("❌ Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterAppointmentServiceServer(grpcServer, appointmentGRPC.NewAppointmentHandler(aApp))
		log.Printf("🚀 gRPC server running at :%s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Failed to serve gRPC: %v", err)
		}
	}()

	// --- Start HTTP Server ---
	log.Printf("✅ HTTP server running at :%s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
