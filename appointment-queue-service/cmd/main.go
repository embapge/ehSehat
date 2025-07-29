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

	// --- gRPC Server ---
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("❌ Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterAppointmentServiceServer(grpcServer, appointmentGRPC.NewAppointmentHandler(aApp))
		log.Println("🚀 gRPC server running at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Failed to serve gRPC: %v", err)
		}
	}()

	// --- Start HTTP Server ---
	log.Println("✅ HTTP server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
