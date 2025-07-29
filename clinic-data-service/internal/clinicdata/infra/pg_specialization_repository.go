package infra

import (
	"context"
	"database/sql"
	"time"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/google/uuid"
)

type pgSpecializationRepository struct {
	db *sql.DB
}

func NewPGSpecializationRepository(db *sql.DB) *pgSpecializationRepository {
	return &pgSpecializationRepository{db}
}

// Create inserts a new specialization into the database
func (r *pgSpecializationRepository) Create(s *domain.Specialization) (*domain.Specialization, error) {
	query := `
		INSERT INTO specializations (
			id, name,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		) VALUES (
			$1, $2,
			$3, $4, $5, $6,
			$7, $8, $9, $10,
			$11, $12
		)
	`

	now := time.Now()
	if s.ID == "" {
		s.ID = uuid.New().String()
	}

	_, err := r.db.ExecContext(context.Background(), query,
		s.ID, s.Name,
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

// GetByID fetches a specialization by ID
func (r *pgSpecializationRepository) GetByID(id string) (*domain.Specialization, error) {
	query := `
		SELECT id, name,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		FROM specializations
		WHERE id = $1
	`

	row := r.db.QueryRowContext(context.Background(), query, id)
	var s domain.Specialization
	err := row.Scan(
		&s.ID, &s.Name,
		&s.CreatedBy, &s.CreatedName, &s.CreatedEmail, &s.CreatedRole,
		&s.UpdatedBy, &s.UpdatedName, &s.UpdatedEmail, &s.UpdatedRole,
		&s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetAll returns all specializations
func (r *pgSpecializationRepository) GetAll() ([]domain.Specialization, error) {
	query := `
		SELECT id, name,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		FROM specializations
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Specialization
	for rows.Next() {
		var s domain.Specialization
		err := rows.Scan(
			&s.ID, &s.Name,
			&s.CreatedBy, &s.CreatedName, &s.CreatedEmail, &s.CreatedRole,
			&s.UpdatedBy, &s.UpdatedName, &s.UpdatedEmail, &s.UpdatedRole,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

// Update modifies an existing specialization
func (r *pgSpecializationRepository) Update(s *domain.Specialization) (*domain.Specialization, error) {
	query := `
		UPDATE specializations SET
			name = $1,
			updated_by = $2, updated_name = $3, updated_email = $4, updated_role = $5,
			updated_at = $6
		WHERE id = $7
	`

	now := time.Now()
	_, err := r.db.ExecContext(context.Background(), query,
		s.Name,
		s.UpdatedBy, s.UpdatedName, s.UpdatedEmail, s.UpdatedRole,
		now, s.ID,
	)
	if err != nil {
		return nil, err
	}
	s.UpdatedAt = now
	return s, nil
}

// Delete removes a specialization by ID
func (r *pgSpecializationRepository) Delete(id string) error {
	query := `DELETE FROM specializations WHERE id = $1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}
