package infra

import (
	"context"
	"database/sql"
	"time"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/google/uuid"
)

type pgDoctorRepository struct {
	db *sql.DB
}

func NewPGDoctorRepository(db *sql.DB) *pgDoctorRepository {
	return &pgDoctorRepository{db}
}

// Create inserts a new doctor into the database
func (r *pgDoctorRepository) Create(d *domain.Doctor) (*domain.Doctor, error) {
	query := `
		INSERT INTO doctors (
			id, user_id, name, email, specialization_id, age, consultation_fee,
			years_of_experience, license_number, phone_number, is_active,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11,
			$12, $13, $14, $15,
			$16, $17, $18, $19,
			$20, $21
		)
	`

	now := time.Now()
	if d.ID == "" {
		d.ID = uuid.New().String()
	}

	_, err := r.db.ExecContext(context.Background(), query,
		d.ID, d.UserID, d.Name, d.Email, d.SpecializationID, d.Age, d.ConsultationFee,
		d.YearsOfExperience, d.LicenseNumber, d.PhoneNumber, d.IsActive,
		d.CreatedBy, d.CreatedName, d.CreatedEmail, d.CreatedRole,
		d.UpdatedBy, d.UpdatedName, d.UpdatedEmail, d.UpdatedRole,
		now, now,
	)
	if err != nil {
		return nil, err
	}

	d.CreatedAt = now
	d.UpdatedAt = now
	return d, nil
}

// GetByID retrieves a doctor by ID
func (r *pgDoctorRepository) GetByID(id string) (*domain.Doctor, error) {
	query := `SELECT 
		id, user_id, name, email, specialization_id, age, consultation_fee,
		years_of_experience, license_number, phone_number, is_active,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM doctors WHERE id = $1`

	row := r.db.QueryRowContext(context.Background(), query, id)
	var d domain.Doctor
	err := row.Scan(
		&d.ID, &d.UserID, &d.Name, &d.Email, &d.SpecializationID, &d.Age, &d.ConsultationFee,
		&d.YearsOfExperience, &d.LicenseNumber, &d.PhoneNumber, &d.IsActive,
		&d.CreatedBy, &d.CreatedName, &d.CreatedEmail, &d.CreatedRole,
		&d.UpdatedBy, &d.UpdatedName, &d.UpdatedEmail, &d.UpdatedRole,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// GetByEmail retrieves a doctor by Email
func (r *pgDoctorRepository) GetByEmail(email string) (*domain.Doctor, error) {
	query := `SELECT 
		id, user_id, name, email, specialization_id, age, consultation_fee,
		years_of_experience, license_number, phone_number, is_active,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM doctors WHERE email = $1 LIMIT 1`

	row := r.db.QueryRowContext(context.Background(), query, email)
	var d domain.Doctor
	err := row.Scan(
		&d.ID, &d.UserID, &d.Name, &d.Email, &d.SpecializationID, &d.Age, &d.ConsultationFee,
		&d.YearsOfExperience, &d.LicenseNumber, &d.PhoneNumber, &d.IsActive,
		&d.CreatedBy, &d.CreatedName, &d.CreatedEmail, &d.CreatedRole,
		&d.UpdatedBy, &d.UpdatedName, &d.UpdatedEmail, &d.UpdatedRole,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// GetAll retrieves all doctors
func (r *pgDoctorRepository) GetAll() ([]domain.Doctor, error) {
	query := `SELECT 
		id, user_id, name, email, specialization_id, age, consultation_fee,
		years_of_experience, license_number, phone_number, is_active,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM doctors ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []domain.Doctor
	for rows.Next() {
		var d domain.Doctor
		err := rows.Scan(
			&d.ID, &d.UserID, &d.Name, &d.Email, &d.SpecializationID, &d.Age, &d.ConsultationFee,
			&d.YearsOfExperience, &d.LicenseNumber, &d.PhoneNumber, &d.IsActive,
			&d.CreatedBy, &d.CreatedName, &d.CreatedEmail, &d.CreatedRole,
			&d.UpdatedBy, &d.UpdatedName, &d.UpdatedEmail, &d.UpdatedRole,
			&d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}
	return doctors, nil
}

// Update modifies an existing doctor
func (r *pgDoctorRepository) Update(d *domain.Doctor) (*domain.Doctor, error) {
	query := `
		UPDATE doctors SET
			name = $1,
			email = $2,
			specialization_id = $3,
			age = $4,
			consultation_fee = $5,
			years_of_experience = $6,
			license_number = $7,
			phone_number = $8,
			is_active = $9,
			updated_by = $10,
			updated_name = $11,
			updated_email = $12,
			updated_role = $13,
			updated_at = $14
		WHERE id = $15
	`

	now := time.Now()
	_, err := r.db.ExecContext(context.Background(), query,
		d.Name, d.Email, d.SpecializationID, d.Age, d.ConsultationFee,
		d.YearsOfExperience, d.LicenseNumber, d.PhoneNumber, d.IsActive,
		d.UpdatedBy, d.UpdatedName, d.UpdatedEmail, d.UpdatedRole,
		now, d.ID,
	)
	if err != nil {
		return nil, err
	}
	d.UpdatedAt = now
	return d, nil
}

// Delete removes a doctor by ID
func (r *pgDoctorRepository) Delete(id string) error {
	query := `DELETE FROM doctors WHERE id = $1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}
