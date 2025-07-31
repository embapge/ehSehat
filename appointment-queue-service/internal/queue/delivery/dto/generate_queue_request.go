package dto

type GenerateQueueRequest struct {
	DoctorID             string  `json:"doctor_id"`
	UserID               string  `json:"user_id"`
	UserName             string  `json:"user_name"`
	UserRole             string  `json:"user_role"`
	UserEmail            string  `json:"user_email"`
	AppointmentID        *uint   `json:"appointment_id"`
	PatientID            *string `json:"patient_id"`
	PatientName          *string `json:"patient_name"`
	DoctorName           string  `json:"doctor_name"`
	DoctorSpecialization string  `json:"doctor_specialization"`
	Type                 string  `json:"type"` // online / offline
}
