package domain

// Implement struct/model for Consultation MongoDB collection
import (
	"time"
)

//   string id = 1; // unique identifier for the consultation
//     string user_id = 2; // unique identifier for the consultation
//     string patient_id = 3; // from master-service
//     string patient_name = 4; // snapshot, optional
//     string doctor_id = 5; // from master-service
//     string doctor_name = 6; // snapshot, optional
//     string room_id = 7; // from master-service
//     string room_name = 8; // snapshot, optional
//     string symptoms = 9;
//     repeated Prescription prescription = 10; // list of prescriptions
//     string diagnosis = 11;
//     string date = 12;
//     google.protobuf.Timestamp created_at = 13;
//     google.protobuf.Timestamp updated_at = 14;

type Consultation struct {
	ID           string         `bson:"_id,omitempty" json:"id"`
	UserID       string         `bson:"user_id" json:"user_id"`
	PatientID    string         `bson:"patient_id" json:"patient_id"`
	PatientName  string         `bson:"patient_name,omitempty" json:"patient_name,omitempty"`
	DoctorID     string         `bson:"doctor_id" json:"doctor_id"`
	DoctorName   string         `bson:"doctor_name,omitempty" json:"doctor_name,omitempty"`
	RoomID       string         `bson:"room_id" json:"room_id"`
	RoomName     string         `bson:"room_name,omitempty" json:"room_name,omitempty"`
	Symptoms     string         `bson:"symptoms" json:"symptoms"`
	Prescription []Prescription `bson:"prescription,omitempty" json:"prescription,omitempty"`
	Diagnosis    string         `bson:"diagnosis" json:"diagnosis"`
	Date         time.Time      `bson:"date" json:"date"`
	CreatedAt    time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `bson:"updated_at" json:"updated_at"`
}

type Prescription struct {
	MedicineName string `bson:"medicine_name" json:"medicine_name"`
	Dose         string `bson:"dose" json:"dose"`
}
