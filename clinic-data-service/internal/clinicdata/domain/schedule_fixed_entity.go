package domain

import "time"

// ScheduleFixed mewakili jadwal tetap dokter mingguan
type ScheduleFixed struct {
	ID        string
	DoctorID  string
	RoomID    string
	DayOfWeek int    // 1 = Monday, ..., 7 = Sunday
	StartTime string // format "HH:MM"
	EndTime   string // format "HH:MM"
	Status    string // "active" / "inactive"

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

// ScheduleFixedRepository interface untuk akses DB
type ScheduleFixedRepository interface {
	Create(s *ScheduleFixed) (*ScheduleFixed, error)
	GetByDoctorID(doctorID string) ([]ScheduleFixed, error)
	Update(s *ScheduleFixed) (*ScheduleFixed, error)
}
