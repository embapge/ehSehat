package domain

import "time"

type QueueModel struct {
	ID                   uint      `gorm:"primaryKey" json:"id,omitempty"`
	UserID               string    `json:"user_id,omitempty"`      // user yang create antrian
	UserName             string    `json:"user_name,omitempty"`    // snapshot
	UserRole             string    `json:"user_role,omitempty"`    // admin/member
	UserEmail            string    `json:"user_email,omitempty"`   // admin/member
	PatientID            *string   `json:"patient_id,omitempty"`   // optional
	PatientName          *string   `json:"patient_name,omitempty"` // optional
	DoctorID             string    `json:"doctor_id,omitempty"`
	DoctorName           string    `json:"doctor_name,omitempty"`
	DoctorSpecialization string    `json:"doctor_specialization,omitempty"`
	AppointmentID        *uint     `json:"appointment_id,omitempty"` // nullable
	Type                 string    `json:"type,omitempty"`           // online / offline
	QueueNumber          int       `json:"queue_number,omitempty"`   // di-generate per dokter/hari
	StartFrom            time.Time `json:"start_from,omitempty"`     // estimasi masuk ruangan
	Status               string    `json:"status,omitempty"`         // active, fail, done
	CreatedAt            time.Time `json:"created_at,omitempty"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
}

func (QueueModel) TableName() string {
	return "queues"
}
