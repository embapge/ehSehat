package domain

import "time"

type QueueModel struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	UserID               uint      `json:"user_id"`      // user yang create antrian
	UserName             string    `json:"user_name"`    // snapshot
	UserRole             string    `json:"user_role"`    // admin/member
	PatientID            *uint     `json:"patient_id"`   // optional
	PatientName          *string   `json:"patient_name"` // optional
	DoctorID             uint      `json:"doctor_id"`
	DoctorName           string    `json:"doctor_name"`
	DoctorSpecialization string    `json:"doctor_specialization"`
	AppointmentID        *uint     `json:"appointment_id"` // nullable
	Type                 string    `json:"type"`           // online / offline
	QueueNumber          int       `json:"queue_number"`   // di-generate per dokter/hari
	StartFrom            time.Time `json:"start_from"`     // estimasi masuk ruangan
	Status               string    `json:"status"`         // active, fail, done
	CreatedAt            time.Time `json:"created_at"`
}
