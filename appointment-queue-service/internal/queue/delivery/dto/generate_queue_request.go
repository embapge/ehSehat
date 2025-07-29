package dto

type GenerateQueueRequest struct {
	DoctorID             uint    `json:"doctor_id"`
	UserID               uint    `json:"user_id"`
	UserName             string  `json:"user_name"`
	UserRole             string  `json:"user_role"`
	AppointmentID        *uint   `json:"appointment_id"`
	PatientID            *uint   `json:"patient_id"`
	PatientName          *string `json:"patient_name"`
	DoctorName           string  `json:"doctor_name"`
	DoctorSpecialization string  `json:"doctor_specialization"`
	Type                 string  `json:"type"` // online / offline
}
