package domain

import "time"

type UserSnapShot struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name,omitempty" json:"name"`
	Email string `bson:"email,omitempty" json:"email"`
	Role  string `bson:"role,omitempty" json:"role"`
}

type PatientSnapshot struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
	Age  int32  `bson:"age,omitempty" json:"age"`
}

type DoctorSnapshot struct {
	ID             string `bson:"_id,omitempty" json:"id"`
	Name           string `bson:"name,omitempty" json:"name"`
	Specialization string `bson:"specialization,omitempty" json:"specialization"`
}

type RoomSnapshot struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}

type CreatedBySnapshot struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name,omitempty" json:"name"`
	Email     string    `bson:"email,omitempty" json:"email"`
	Role      string    `bson:"role,omitempty" json:"role"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
}

type UpdatedBySnapshot struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name,omitempty" json:"name"`
	Email     string    `bson:"email,omitempty" json:"email"`
	Role      string    `bson:"role,omitempty" json:"role"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
}
