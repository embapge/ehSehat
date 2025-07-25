package main

import (
	"log"
	"net/http"

	"appointment-queue-service/config"

	// Appointment layers
	appointmentApp "appointment-queue-service/internal/appointment/app"
	appointmentHTTP "appointment-queue-service/internal/appointment/delivery/http"
	appointmentDomain "appointment-queue-service/internal/appointment/domain"
	appointmentRepo "appointment-queue-service/internal/appointment/domain"

	// Queue layers
	queueApp "appointment-queue-service/internal/queue/app"
	queueHTTP "appointment-queue-service/internal/queue/delivery/http"
	"appointment-queue-service/internal/queue/domain"
	queueRepo "appointment-queue-service/internal/queue/domain"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Load .env or fallback to OS environment
	config.LoadEnv()

	// Initialize DB
	db := config.DBInit()

	// -------------------- Queue Setup --------------------
	qRepo := queueRepo.NewQueueRepository(db)
	qService := domain.NewQueueService(qRepo)
	qApp := queueApp.NewQueueApp(qService)

	// -------------------- Appointment Setup --------------------
	aRepo := appointmentRepo.NewAppointmentRepository(db)
	aService := appointmentDomain.NewAppointmentService(aRepo)
	aApp := appointmentApp.NewAppointmentApp(aService)

	// Router setup
	router := httprouter.New()

	// Register handlers to routes
	queueHTTP.NewQueueHandler(router, qApp)
	appointmentHTTP.NewAppointmentHandler(router, aApp)

	// Run HTTP server
	log.Println("✅ Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
