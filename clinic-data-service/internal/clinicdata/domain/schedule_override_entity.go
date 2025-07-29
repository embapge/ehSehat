package domain

import "time"

// ScheduleOverride represents the override schedule for a doctor
type ScheduleOverride struct {
	ID        string
	DoctorID  string
	RoomID    string
	DayOfWeek int
	StartTime string
	EndTime   string
	Status    string

	CreatedBy    *string
	CreatedName  string
	CreatedEmail string
	CreatedRole  string
	CreatedAt    time.Time

	UpdatedBy    *string
	UpdatedName  string
	UpdatedEmail string
	UpdatedRole  string
	UpdatedAt    time.Time
}

// ScheduleOverrideRepository defines DB behavior for ScheduleOverride
type ScheduleOverrideRepository interface {
	Create(override *ScheduleOverride) (*ScheduleOverride, error)
	GetByDoctorID(doctorID string) ([]ScheduleOverride, error)
	Update(override *ScheduleOverride) (*ScheduleOverride, error)
	Delete(id string) error
}
