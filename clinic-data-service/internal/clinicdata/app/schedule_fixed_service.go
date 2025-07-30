package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"log"
	"strings"
)

// ScheduleFixedService interface defines business logic for fixed schedule
type ScheduleFixedService interface {
	Create(s *domain.ScheduleFixed) (*domain.ScheduleFixed, error)
	GetByDoctorID(doctorID string) ([]domain.ScheduleFixed, error)
	Update(s *domain.ScheduleFixed) (*domain.ScheduleFixed, error)
}

// scheduleFixedService implements ScheduleFixedService
type scheduleFixedService struct {
	repo domain.ScheduleFixedRepository
}

// NewScheduleFixedService returns a new service instance
func NewScheduleFixedService(r domain.ScheduleFixedRepository) ScheduleFixedService {
	return &scheduleFixedService{repo: r}
}

// Create inserts a new fixed schedule with validation
func (s *scheduleFixedService) Create(sf *domain.ScheduleFixed) (*domain.ScheduleFixed, error) {
	log.Printf("DEBUG create validation: DoctorID='%s' RoomID='%s' DayOfWeek=%d Start='%s' End='%s'",
		sf.DoctorID, sf.RoomID, sf.DayOfWeek, sf.StartTime, sf.EndTime)

	// Validasi field wajib
	if strings.TrimSpace(sf.DoctorID) == "" ||
		strings.TrimSpace(sf.RoomID) == "" ||
		sf.DayOfWeek < 1 || sf.DayOfWeek > 7 ||
		strings.TrimSpace(sf.StartTime) == "" ||
		strings.TrimSpace(sf.EndTime) == "" ||
		strings.TrimSpace(sf.Status) == "" {
		return nil, ErrMissingFields
	}

	return s.repo.Create(sf)
}

// GetByDoctorID returns all fixed schedules by doctor
func (s *scheduleFixedService) GetByDoctorID(doctorID string) ([]domain.ScheduleFixed, error) {
	if strings.TrimSpace(doctorID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.GetByDoctorID(doctorID)
}

// Update modifies an existing fixed schedule
func (s *scheduleFixedService) Update(sf *domain.ScheduleFixed) (*domain.ScheduleFixed, error) {
	if strings.TrimSpace(sf.ID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.Update(sf)
}
