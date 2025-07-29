package infra

import (
	"clinic-data-service/internal/clinicdata/domain"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// pgScheduleOverrideRepo implements ScheduleOverrideRepository
type pgScheduleOverrideRepo struct {
	db *sql.DB
}

// NewPGScheduleOverrideRepository returns new repository instance
func NewPGScheduleOverrideRepository(db *sql.DB) domain.ScheduleOverrideRepository {
	return &pgScheduleOverrideRepo{db: db}
}

// Create inserts new override schedule
func (r *pgScheduleOverrideRepo) Create(o *domain.ScheduleOverride) (*domain.ScheduleOverride, error) {
	id := uuid.New().String()
	now := time.Now()

	query := `
		INSERT INTO schedule_overrides (
			id, doctor_id, room_id, day_of_week, start_time, end_time, status,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,
				$8,$9,$10,$11,
				$12,$13,$14,$15,
				$16,$17)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		id, o.DoctorID, o.RoomID, o.DayOfWeek, o.StartTime, o.EndTime, o.Status,
		o.CreatedBy, o.CreatedName, o.CreatedEmail, o.CreatedRole,
		o.UpdatedBy, o.UpdatedName, o.UpdatedEmail, o.UpdatedRole,
		now, now,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	o.ID = id
	o.CreatedAt = now
	o.UpdatedAt = now
	return o, nil
}

// GetByDoctorID returns overrides by doctor ID
func (r *pgScheduleOverrideRepo) GetByDoctorID(doctorID string) ([]domain.ScheduleOverride, error) {
	query := `
		SELECT id, doctor_id, room_id, day_of_week, start_time, end_time, status,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		FROM schedule_overrides
		WHERE doctor_id = $1
		ORDER BY day_of_week ASC, start_time ASC
	`

	rows, err := r.db.Query(query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.ScheduleOverride
	for rows.Next() {
		var o domain.ScheduleOverride
		err := rows.Scan(
			&o.ID, &o.DoctorID, &o.RoomID, &o.DayOfWeek, &o.StartTime, &o.EndTime, &o.Status,
			&o.CreatedBy, &o.CreatedName, &o.CreatedEmail, &o.CreatedRole,
			&o.UpdatedBy, &o.UpdatedName, &o.UpdatedEmail, &o.UpdatedRole,
			&o.CreatedAt, &o.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, o)
	}
	return result, nil
}

// Update modifies schedule override
func (r *pgScheduleOverrideRepo) Update(o *domain.ScheduleOverride) (*domain.ScheduleOverride, error) {
	now := time.Now()

	query := `
		UPDATE schedule_overrides
		SET room_id = $1, start_time = $2, end_time = $3, status = $4,
			updated_by = $5, updated_name = $6, updated_email = $7, updated_role = $8,
			updated_at = $9
		WHERE id = $10
	`

	_, err := r.db.Exec(
		query,
		o.RoomID, o.StartTime, o.EndTime, o.Status,
		o.UpdatedBy, o.UpdatedName, o.UpdatedEmail, o.UpdatedRole,
		now, o.ID,
	)
	if err != nil {
		return nil, err
	}

	o.UpdatedAt = now
	return o, nil
}

// Delete removes override by ID
func (r *pgScheduleOverrideRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM schedule_overrides WHERE id = $1`, id)
	return err
}
