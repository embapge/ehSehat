package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"clinic-data-service/config"
	"clinic-data-service/internal/clinicdata/app"
	grpcHandler "clinic-data-service/internal/clinicdata/delivery/grpc"
	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	"clinic-data-service/internal/clinicdata/infra"
	"ehSehat/libs/utils/rabbitmqown"
)

func main() {
	// STEP 1: Load ENV
	env := config.LoadEnv()

	// STEP 2: Init DB
	db := infra.InitDB(env)

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

	// STEP 3: Init Repositories (infra layer)
	patientRepo := infra.NewPGPatientRepository(db)
	doctorRepo := infra.NewPGDoctorRepository(db)
	specRepo := infra.NewPGSpecializationRepository(db)
	roomRepo := infra.NewPGRoomRepository(db)
	scheduleFixedRepo := infra.NewPGScheduleFixedRepository(db)
	scheduleOverrideRepo := infra.NewPGScheduleOverrideRepository(db)

	// STEP 4: Init Services (app layer)
	patientService := app.NewPatientService(patientRepo, ch)
	doctorService := app.NewDoctorService(doctorRepo, ch)
	specService := app.NewSpecializationService(specRepo)
	roomService := app.NewRoomService(roomRepo)
	scheduleFixedService := app.NewScheduleFixedService(scheduleFixedRepo)
	scheduleOverrideService := app.NewScheduleOverrideService(scheduleOverrideRepo)

	// STEP 5: Init Handler (delivery layer)
	handler := grpcHandler.NewGRPCHandler(
		patientService,
		doctorService,
		specService,
		roomService,
		scheduleFixedService,
		scheduleOverrideService,
	)

	// STEP 6: Setup gRPC server
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50052"
	}
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}

	grpcServer := grpc.NewServer()
	clinicdatapb.RegisterClinicDataServiceServer(grpcServer, handler)

	log.Printf("gRPC server running on %s", addr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
