# Create Appointment
curl --location 'http://localhost:8080/appointments' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "user_full_name": "Bagus",
    "doctor_id": 2,
    "doctor_name": "Dr. Siti",
    "doctor_specialization": "Umum",
    "appointment_at": "2025-08-01T09:00:00Z",
    "is_paid": false,
    "status": "unpaid"
  }'

# Get All Data Appointment
curl http://localhost:8080/appointments

# Get Find By ID Appointment
curl --location 'http://localhost:8080/appointments/:id'

# Get Find By User ID Appointment
curl --location 'http://localhost:8080/appointments-by-user/:user_id'

# Patch Appointment
curl --location --request PATCH 'http://localhost:8080/appointments/1' \
--header 'Content-Type: application/json' \
--data '{
    "appointment_at": "2025-08-02T10:30:00Z",
    "status": "rescheduled"
  }'

# Update Status Paid Appointment
curl --location --request PUT 'http://localhost:8080/appointments/1/mark-paid'

# Create Queue
curl --location 'http://localhost:8080/queues' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 2,
    "user_name": "Baguss",
    "user_role": "patient",
    "patient_id": 2,
    "doctor_id": 10,
    "doctor_name": "Dr. Smith",
    "doctor_specialization": "Dermatology",
    "appointment_id": 2,
    "type": "offline"
  }'

# Get Queue by ID
curl -X GET http://localhost:8080/queues/1

# Get Today's Queues by Doctor
curl -X GET http://localhost:8080/queues-today/10

# Update Queue
curl -X PUT http://localhost:8080/queues/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "user_id": 1,
    "user_name": "admin",
    "user_role": "admin",
    "patient_id": 2,
    "patient_name": "Budi Update",
    "doctor_id": 10,
    "doctor_name": "Dr. Smith",
    "doctor_specialization": "Dermatology",
    "appointment_id": 5,
    "type": "offline",
    "queue_number": 3,
    "start_from": "2025-07-29T09:30:00Z",
    "status": "done"
  }'

# Generate Next Queue (auto-generate queue number)
curl -X POST http://localhost:8080/queues/generate \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "user_name": "admin",
    "user_role": "admin",
    "doctor_id": 10,
    "doctor_name": "Dr. Smith",
    "doctor_specialization": "Dermatology",
    "appointment_id": 5,
    "patient_id": 2,
    "patient_name": "Budi",
    "type": "online"
  }'
