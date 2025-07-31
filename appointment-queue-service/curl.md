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