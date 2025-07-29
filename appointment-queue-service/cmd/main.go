package main

import (
	"log"
	"net"
	"net/http"

	"appointment-queue-service/config"

	// Queue layers
	queueApp "appointment-queue-service/internal/queue/app"
	queueHTTP "appointment-queue-service/internal/queue/delivery/http"
	queueDomain "appointment-queue-service/internal/queue/domain"
	queueRepo "appointment-queue-service/internal/queue/domain"

	// Appointment layers
	appointmentApp "appointment-queue-service/internal/appointment/app"
	appointmentGRPC "appointment-queue-service/internal/appointment/delivery/grpc"
	"appointment-queue-service/internal/appointment/delivery/grpc/pb"
	appointmentHTTP "appointment-queue-service/internal/appointment/delivery/http"
	appointmentDomain "appointment-queue-service/internal/appointment/domain"
	appointmentRepo "appointment-queue-service/internal/appointment/domain"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

func main() {
	// Load .env or fallback to OS environment
	config.LoadEnv()

	// Initialize DB
	db := config.DBInit()

	// -------------------- Queue Setup --------------------
	qRepo := queueRepo.NewQueueRepository(db)
	qService := queueDomain.NewQueueService(qRepo)
	qApp := queueApp.NewQueueApp(qService)

	// -------------------- Appointment Setup --------------------
	aRepo := appointmentRepo.NewAppointmentRepository(db)
	aService := appointmentDomain.NewAppointmentService(aRepo)
	aApp := appointmentApp.NewAppointmentApp(aService)

	// -------------------- HTTP Setup --------------------
	router := httprouter.New()
	queueHTTP.NewQueueHandler(router, qApp)
	appointmentHTTP.NewAppointmentHandler(router, aApp)

	// -------------------- gRPC Setup --------------------
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("❌ Failed to listen on port 50051: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcHandler := appointmentGRPC.NewAppointmentHandler(aApp)
		pb.RegisterAppointmentServiceServer(grpcServer, grpcHandler)

		log.Println("✅ gRPC server running at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Failed to serve gRPC: %v", err)
		}
	}()

	// -------------------- Start HTTP --------------------
	log.Println("✅ HTTP server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
