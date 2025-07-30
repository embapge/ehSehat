package infra

import (
	"context"
	"database/sql"
	"time"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/google/uuid"
)

type pgScheduleFixedRepository struct {
	db *sql.DB
}

func NewPGScheduleFixedRepository(db *sql.DB) *pgScheduleFixedRepository {
	return &pgScheduleFixedRepository{db}
}

// Create inserts a new fixed schedule into the database
func (r *pgScheduleFixedRepository) Create(s *domain.ScheduleFixed) (*domain.ScheduleFixed, error) {
	query := `
		INSERT INTO schedule_fixed (
			id, doctor_id, room_id, day_of_week, start_time, end_time, status,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11,
			$12, $13, $14, $15,
			$16, $17
		)
	`

	now := time.Now()
	if s.ID == "" {
		s.ID = uuid.New().String()
	}

	_, err := r.db.ExecContext(context.Background(), query,
		s.ID, s.DoctorID, s.RoomID, s.DayOfWeek, s.StartTime, s.EndTime, s.Status,
		s.CreatedBy, s.CreatedName, s.CreatedEmail, s.CreatedRole,
		s.UpdatedBy, s.UpdatedName, s.UpdatedEmail, s.UpdatedRole,
		now, now,
	)
	if err != nil {
		return nil, err
	}

	s.CreatedAt = now
	s.UpdatedAt = now
	return s, nil
}

// GetByDoctorID retrieves all fixed schedules for a specific doctor
func (r *pgScheduleFixedRepository) GetByDoctorID(doctorID string) ([]domain.ScheduleFixed, error) {
	query := `SELECT 
		id, doctor_id, room_id, day_of_week, start_time, end_time, status,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM schedule_fixed
	WHERE doctor_id = $1
	ORDER BY day_of_week ASC, start_time ASC`

	rows, err := r.db.QueryContext(context.Background(), query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []domain.ScheduleFixed
	for rows.Next() {
		var s domain.ScheduleFixed
		err := rows.Scan(
			&s.ID, &s.DoctorID, &s.RoomID, &s.DayOfWeek, &s.StartTime, &s.EndTime, &s.Status,
			&s.CreatedBy, &s.CreatedName, &s.CreatedEmail, &s.CreatedRole,
			&s.UpdatedBy, &s.UpdatedName, &s.UpdatedEmail, &s.UpdatedRole,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

// Update modifies an existing fixed schedule
func (r *pgScheduleFixedRepository) Update(s *domain.ScheduleFixed) (*domain.ScheduleFixed, error) {
	query := `
		UPDATE schedule_fixed SET
			room_id = $1,
			day_of_week = $2,
			start_time = $3,
			end_time = $4,
			status = $5,
			updated_by = $6,
			updated_name = $7,
			updated_email = $8,
			updated_role = $9,
			updated_at = $10
		WHERE id = $11
	`

	now := time.Now()
	_, err := r.db.ExecContext(context.Background(), query,
		s.RoomID, s.DayOfWeek, s.StartTime, s.EndTime, s.Status,
		s.UpdatedBy, s.UpdatedName, s.UpdatedEmail, s.UpdatedRole,
		now, s.ID,
	)
	if err != nil {
		return nil, err
	}

	s.UpdatedAt = now
	return s, nil
}
