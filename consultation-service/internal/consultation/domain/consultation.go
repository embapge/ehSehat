package domain

// Implement struct/model for Consultation MongoDB collection
import (
	"time"
)

type Consultation struct {
	ID            string            `bson:"_id,omitempty" json:"id"`
	QueueID       string            `bson:"queue_id,omitempty" json:"queue_id"`
	AppointmentID string            `bson:"appointment_id,omitempty" json:"appointment_id"`
	User          UserSnapShot      `bson:"user,omitempty" json:"user"`
	Patient       PatientSnapshot   `bson:"patient,omitempty" json:"patient"`
	Doctor        DoctorSnapshot    `bson:"doctor,omitempty" json:"doctor"`
	Room          RoomSnapshot      `bson:"room,omitempty" json:"room"`
	Symptoms      string            `bson:"symptoms" json:"symptoms"`
	Prescription  []Prescription    `bson:"prescription,omitempty" json:"prescription,omitempty"`
	Diagnosis     string            `bson:"diagnosis" json:"diagnosis"`
	Date          time.Time         `bson:"date" json:"date"`
	Status        string            `bson:"status" json:"status"`
	Amount        float64           `bson:"amount" json:"amount"`               // Amount for the consultation
	TotalPayment  float64           `bson:"total_payment" json:"total_payment"` // Total payment amount
	CreatedBy     CreatedBySnapshot `bson:"created_by,omitempty" json:"created_by"`
	CreatedAt     time.Time         `bson:"created_at" json:"created_at"`
	UpdatedBy     UpdatedBySnapshot `bson:"updated_by,omitempty" json:"updated_by"`
	UpdatedAt     time.Time         `bson:"updated_at" json:"updated_at"`
}

type Prescription struct {
	MedicineName string `bson:"medicine_name" json:"medicine_name"`
	Dose         string `bson:"dose" json:"dose"`
}
