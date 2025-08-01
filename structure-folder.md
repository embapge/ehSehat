```
└── 📁ehSehat
    └── 📁.github
        └── 📁workflows
            ├── deploy.yml
    └── 📁appointment-queue-service
        └── 📁cmd
            ├── main.go
        └── 📁config
            ├── database.go
        └── 📁docs
            ├── docs.go
            ├── swagger.json
            ├── swagger.yaml
        └── 📁internal
            └── 📁appointment
                └── 📁app
                    ├── appointment_app.go
                └── 📁delivery
                    └── 📁grpc
                        └── 📁pb
                            ├── appointment_grpc.pb.go
                            ├── appointment.pb.go
                        ├── handler.go
                    └── 📁http
                        ├── appointment_handler.go
                └── 📁domain
                    ├── appointment_model.go
                    ├── appointment_repository.go
                    ├── appointment_service.go
                └── 📁proto
                    ├── appointment.proto
                └── 📁unitTest
                    ├── appointment_service_test.go
            └── 📁queue
                └── 📁app
                    ├── queue_app.go
                └── 📁delivery
                    └── 📁dto
                        ├── generate_queue_request.go
                    └── 📁grpc
                        └── 📁pb
                    └── 📁http
                        ├── queue_handler.go
                └── 📁domain
                    ├── queue_model.go
                    ├── queue_repository.go
                    ├── queue_service.go
                └── 📁proto
                    ├── queue.proto
                └── 📁unitTest
                    ├── queue_service_test.go
        ├── .env
        ├── curl.md
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
        ├── query.sql
    └── 📁auth-service
        └── 📁cmd
            ├── main.go
        └── 📁docs
            ├── docs.go
            ├── swagger.json
            ├── swagger.yaml
        └── 📁internal
            └── 📁auth
                └── 📁app
                    ├── auth_app.go
                    ├── error.go
                    ├── jwt_manager.go
                    ├── password_hasher.go
                └── 📁config
                    ├── db.go
                └── 📁delivery
                    └── 📁grpc
                        └── 📁pb
                            ├── auth_grpc.pb.go
                            ├── auth.pb.go
                        ├── auth_handler.go
                        ├── auth_interceptor.go
                    └── 📁listener
                        ├── doctor_created_listener.go
                        ├── patient_created_listener.go
                └── 📁domain
                    ├── user_repository.go
                    ├── user_service.go
                    ├── user.go
                └── 📁infra
                    ├── mysql_user_repository.go
                └── 📁unitTest
                    ├── auth_handler_test.go
        └── 📁pkg
            └── 📁hasher
                ├── password_hasher_bcrypt.go
            └── 📁jwt
                ├── jwt_manager_standard.go
        ├── .env
        ├── auth.sql
        ├── Best Practice.png
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
    └── 📁clinic-data-service
        └── 📁cmd
            ├── main.go
        └── 📁config
            ├── env.go
        └── 📁internal
            └── 📁clinicdata
                └── 📁app
                    ├── doctor_service.go
                    ├── error.go
                    ├── patient_service.go
                    ├── room_service.go
                    ├── schedule_fixed_service.go
                    ├── schedule_override_service.go
                    ├── specialization_service.go
                └── 📁delivery
                    └── 📁grpc
                        └── 📁clinicdatapb
                            ├── clinic_data_grpc.pb.go
                            ├── clinic_data.pb.go
                        └── 📁utils
                            ├── audit.go
                        ├── doctor_handler.go
                        ├── error.go
                        ├── grpc_handler.go
                        ├── patient_handler.go
                        ├── room_handler.go
                        ├── schedule_fixed_handler.go
                        ├── schedule_override_handler.go
                        ├── specialization_handler.go
                └── 📁domain
                    ├── doctor_entity.go
                    ├── patient_entity.go
                    ├── room_entity.go
                    ├── schedule_fixed_entity.go
                    ├── schedule_override_entity.go
                    ├── specialization_entity.go
                └── 📁infra
                    ├── db.go
                    ├── pg_doctor_repository.go
                    ├── pg_patient_repository.go
                    ├── pg_room_repository.go
                    ├── pg_schedule_fixed_repository.go
                    ├── pg_schedule_override_repository.go
                    ├── pg_specialization_repository.go
        └── 📁migration
            ├── init_schema.sql
        └── 📁proto
            ├── clinic_data_grpc.pb.go
            ├── clinic_data.pb.go
            ├── clinic_data.proto
        ├── .env
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
    └── 📁consultation-service
        └── 📁cmd
            ├── main.go
        └── 📁config
            ├── database.go
        └── 📁internal
            └── 📁consultation
                └── 📁app
                    ├── consultation_app.go
                └── 📁delivery
                    └── 📁grpc
                        └── 📁pb
                            ├── consultation_grpc.pb.go
                            ├── consultation.pb.go
                        ├── consultation_handler.go
                └── 📁domain
                    ├── consultation_repository.go
                    ├── consultation_service.go
                    ├── consultation.go
                    ├── snapshot.go
                └── 📁infra
                    ├── mongo_consultation.go
            └── 📁payment
                └── 📁app
                    ├── payment_app.go
                └── 📁delivery
                    └── 📁grpc
                        └── 📁pb
                            ├── payment_grpc.pb.go
                            ├── payment.pb.go
                        ├── payment_handler.go
                └── 📁domain
                    ├── payment_gateway.go
                    ├── payment_repository.go
                    ├── payment_service.go
                    ├── payment.go
                └── 📁infra
                    ├── mysql_payment.go
                    ├── xendit.go
        ├── .env
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
        ├── mysql_payment.sql
        ├── param.json
    └── 📁diagrams
        └── 📁drawio
            ├── flow-chart-buat-appointment.drawio
            ├── Use Case.drawio
        └── 📁use-case
            ├── use-case-18-July-2025.drawio.png
    └── 📁gateway-service
        └── 📁cmd
            ├── main.go
        └── 📁config
            ├── grpc_client.go
        └── 📁docs
            ├── docs.go
            ├── swagger.json
            ├── swagger.yaml
        └── 📁handler
            └── 📁pb
                ├── appointment_grpc.pb.go
                ├── appointment.pb.go
                ├── auth_grpc.pb.go
                ├── auth.pb.go
                ├── clinic_data_grpc.pb.go
                ├── clinic_data.pb.go
                ├── consultation_grpc.pb.go
                ├── consultation.pb.go
                ├── payment_grpc.pb.go
                ├── payment.pb.go
            ├── gateway_handler.go
        └── 📁middleware
            ├── access_middleware.go
            ├── jwt_middleware.go
        ├── .env
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
    └── 📁libs
        └── 📁utils
            └── 📁grpcmetadata
                ├── metadata.go
            └── 📁rabbitmqown
                ├── payload.go
                ├── rabbitmq.go
            └── 📁security
                ├── jwt_payload.go
                ├── jwt_utils.go
            ├── custom_error.go
            ├── grpc_error.go
            ├── time.go
    └── 📁notification-service
        └── 📁cmd
            ├── main.go
        └── 📁config
            ├── database.go
        └── 📁internal
            └── 📁notification
                └── 📁app
                    ├── notification_app.go
                └── 📁delivery
                    └── 📁listener
                        ├── consultation_listener.go
                └── 📁domain
                    ├── notification_repository.go
                    ├── notification_service.go
                    ├── notification.go
                └── 📁infra
                    ├── mysql_notification.go
        ├── .env
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
        ├── teman_sehat_notifications.sql
    └── 📁proto
        └── 📁appointment
            ├── appointment.proto
        └── 📁auth
            ├── auth.proto
        └── 📁clinic-data-service
            ├── clinic_data.proto
        └── 📁consultation
            ├── consultation.proto
            ├── payment.proto
    ├── docker-compose.yml
    ├── go.mod
    └── go.sum
```
