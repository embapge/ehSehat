package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"log"
	"strings"
)

// ScheduleOverrideService defines the use-case logic
type ScheduleOverrideService interface {
	Create(override *domain.ScheduleOverride) (*domain.ScheduleOverride, error)
	GetByDoctorID(doctorID string) ([]domain.ScheduleOverride, error)
	Update(override *domain.ScheduleOverride) (*domain.ScheduleOverride, error)
	Delete(id string) error
}

// service implementation
type scheduleOverrideService struct {
	repo domain.ScheduleOverrideRepository
}

// Constructor
func NewScheduleOverrideService(r domain.ScheduleOverrideRepository) ScheduleOverrideService {
	return &scheduleOverrideService{repo: r}
}

// Create override schedule with validation
func (s *scheduleOverrideService) Create(o *domain.ScheduleOverride) (*domain.ScheduleOverride, error) {
	log.Printf("DEBUG Create Override: DoctorID='%s' RoomID='%s' Day=%d Start='%s' End='%s' Status='%s'",
		o.DoctorID, o.RoomID, o.DayOfWeek, o.StartTime, o.EndTime, o.Status)

	// Validasi field wajib
	if strings.TrimSpace(o.DoctorID) == "" ||
		strings.TrimSpace(o.RoomID) == "" ||
		o.DayOfWeek < 0 || o.DayOfWeek > 6 ||
		strings.TrimSpace(o.StartTime) == "" ||
		strings.TrimSpace(o.EndTime) == "" ||
		strings.TrimSpace(o.Status) == "" {
		return nil, ErrMissingFields
	}

	return s.repo.Create(o)
}

// GetByDoctorID returns list of overrides for a doctor
func (s *scheduleOverrideService) GetByDoctorID(doctorID string) ([]domain.ScheduleOverride, error) {
	if strings.TrimSpace(doctorID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.GetByDoctorID(doctorID)
}

// Update an override schedule
func (s *scheduleOverrideService) Update(o *domain.ScheduleOverride) (*domain.ScheduleOverride, error) {
	if strings.TrimSpace(o.ID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.Update(o)
}

// Delete an override by ID
func (s *scheduleOverrideService) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingID
	}
	return s.repo.Delete(id)
}
