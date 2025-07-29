package infra

import (
	"context"
	"database/sql"
	"time"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/google/uuid"
)

type pgPatientRepository struct {
	db *sql.DB
}

func NewPGPatientRepository(db *sql.DB) *pgPatientRepository {
	return &pgPatientRepository{db}
}

// Create inserts a new patient into the database
func (r *pgPatientRepository) Create(patient *domain.Patient) (*domain.Patient, error) {
	query := `
		INSERT INTO patients (
			id, user_id, name, email, birth_date, gender, phone_number, address,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12,
			$13, $14, $15, $16,
			$17, $18
		)
	`

	now := time.Now()
	if patient.ID == "" {
		patient.ID = uuid.New().String()
	}
	_, err := r.db.ExecContext(context.Background(), query,
		patient.ID, patient.UserID, patient.Name, patient.Email, patient.BirthDate, patient.Gender,
		patient.PhoneNumber, patient.Address,
		patient.CreatedBy, patient.CreatedName, patient.CreatedEmail, patient.CreatedRole,
		patient.UpdatedBy, patient.UpdatedName, patient.UpdatedEmail, patient.UpdatedRole,
		now, now,
	)
	if err != nil {
		return nil, err
	}

	patient.CreatedAt = now
	patient.UpdatedAt = now
	return patient, nil
}

// GetByID retrieves a patient by ID
func (r *pgPatientRepository) GetByID(id string) (*domain.Patient, error) {
	query := `SELECT 
		id, user_id, name, email, birth_date, gender, phone_number, address,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM patients WHERE id = $1`

	row := r.db.QueryRowContext(context.Background(), query, id)
	var p domain.Patient
	err := row.Scan(
		&p.ID, &p.UserID, &p.Name, &p.Email, &p.BirthDate, &p.Gender, &p.PhoneNumber, &p.Address,
		&p.CreatedBy, &p.CreatedName, &p.CreatedEmail, &p.CreatedRole,
		&p.UpdatedBy, &p.UpdatedName, &p.UpdatedEmail, &p.UpdatedRole,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetByEmail retrieves a patient by Email
func (r *pgPatientRepository) GetByEmail(email string) (*domain.Patient, error) {
	query := `SELECT 
		id, user_id, name, email, birth_date, gender, phone_number, address,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM patients WHERE email = $1 LIMIT 1`

	row := r.db.QueryRowContext(context.Background(), query, email)

	var p domain.Patient
	err := row.Scan(
		&p.ID, &p.UserID, &p.Name, &p.Email, &p.BirthDate, &p.Gender, &p.PhoneNumber, &p.Address,
		&p.CreatedBy, &p.CreatedName, &p.CreatedEmail, &p.CreatedRole,
		&p.UpdatedBy, &p.UpdatedName, &p.UpdatedEmail, &p.UpdatedRole,
		&p.CreatedAt, &p.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Tidak ditemukan
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// GetAll retrieves all patients
func (r *pgPatientRepository) GetAll() ([]domain.Patient, error) {
	query := `SELECT 
		id, user_id, name, email, birth_date, gender, phone_number, address,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM patients ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var p domain.Patient
		err := rows.Scan(
			&p.ID, &p.UserID, &p.Name, &p.Email, &p.BirthDate, &p.Gender, &p.PhoneNumber, &p.Address,
			&p.CreatedBy, &p.CreatedName, &p.CreatedEmail, &p.CreatedRole,
			&p.UpdatedBy, &p.UpdatedName, &p.UpdatedEmail, &p.UpdatedRole,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}

// Update modifies an existing patient
func (r *pgPatientRepository) Update(patient *domain.Patient) (*domain.Patient, error) {
	query := `
		UPDATE patients SET
			name = $1,
			email = $2,
			birth_date = $3,
			gender = $4,
			phone_number = $5,
			address = $6,
			updated_by = $7,
			updated_name = $8,
			updated_email = $9,
			updated_role = $10,
			updated_at = $11
		WHERE id = $12
	`
	now := time.Now()
	_, err := r.db.ExecContext(context.Background(), query,
		patient.Name, patient.Email, patient.BirthDate, patient.Gender,
		patient.PhoneNumber, patient.Address,
		patient.UpdatedBy, patient.UpdatedName, patient.UpdatedEmail, patient.UpdatedRole,
		now, patient.ID,
	)
	if err != nil {
		return nil, err
	}
	patient.UpdatedAt = now
	return patient, nil
}

// Delete removes a patient by ID
func (r *pgPatientRepository) Delete(id string) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err
}
