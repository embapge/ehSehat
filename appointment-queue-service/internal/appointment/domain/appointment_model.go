package domain

import "time"

// type AppointmentModel struct {
// 	ID           uint   `gorm:"primaryKey" json:"id"`
// 	UserID       uint   `json:"user_id"`        // yang booking
// 	UserFullName string `json:"user_full_name"` // snapshot
// 	// PatientID            uint      `json:"patient_id"`   // pasien yang akan berobat
// 	// PatientName          string    `json:"patient_name"` // snapshot
// 	DoctorID             uint      `json:"doctor_id"`
// 	DoctorName           string    `json:"doctor_name"`           // snapshot
// 	DoctorSpecialization string    `json:"doctor_specialization"` // snapshot
// 	AppointmentAt        time.Time `json:"appointment_at"`        // waktu janji temu (start_from)
// 	IsPaid               bool      `json:"is_paid"`               // sudah dibayar atau belum
// 	Status               string    `json:"status"`                // paid, unpaid, void/cancel
// 	CreatedAt            time.Time `json:"created_at"`
// }

type AppointmentModel struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	UserID               uint      `json:"user_id"`        // yang booking
	UserFullName         string    `json:"user_full_name"` // snapshot
	DoctorID             uint      `json:"doctor_id"`
	DoctorName           string    `json:"doctor_name"`           // snapshot
	DoctorSpecialization string    `json:"doctor_specialization"` // snapshot
	AppointmentAt        time.Time `json:"appointment_at"`        // waktu janji temu
	IsPaid               bool      `json:"is_paid"`               // apakah sudah dibayar
	Status               string    `json:"status"`                // paid, unpaid, void/cancel
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (AppointmentModel) TableName() string {
	return "appointments"
}
