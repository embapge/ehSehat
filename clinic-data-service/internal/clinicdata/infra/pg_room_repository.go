package infra

import (
	"context"
	"database/sql"
	"time"

	"clinic-data-service/internal/clinicdata/domain"
	"github.com/google/uuid"
)

type pgRoomRepository struct {
	db *sql.DB
}

func NewPGRoomRepository(db *sql.DB) *pgRoomRepository {
	return &pgRoomRepository{db}
}

// Create inserts a new room into the database
func (r *pgRoomRepository) Create(room *domain.Room) (*domain.Room, error) {
	query := `
		INSERT INTO rooms (
			id, name, is_active,
			created_by, created_name, created_email, created_role,
			updated_by, updated_name, updated_email, updated_role,
			created_at, updated_at
		) VALUES (
			$1, $2, $3,
			$4, $5, $6, $7,
			$8, $9, $10, $11,
			$12, $13
		)
	`

	now := time.Now()
	if room.ID == "" {
		room.ID = uuid.New().String()
	}

	_, err := r.db.ExecContext(context.Background(), query,
		room.ID, room.Name, room.IsActive,
		room.CreatedBy, room.CreatedName, room.CreatedEmail, room.CreatedRole,
		room.UpdatedBy, room.UpdatedName, room.UpdatedEmail, room.UpdatedRole,
		now, now,
	)
	if err != nil {
		return nil, err
	}

	room.CreatedAt = now
	room.UpdatedAt = now
	return room, nil
}

// GetByID retrieves a room by ID
func (r *pgRoomRepository) GetByID(id string) (*domain.Room, error) {
	query := `SELECT 
		id, name, is_active,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM rooms WHERE id = $1`

	row := r.db.QueryRowContext(context.Background(), query, id)
	var room domain.Room
	err := row.Scan(
		&room.ID, &room.Name, &room.IsActive,
		&room.CreatedBy, &room.CreatedName, &room.CreatedEmail, &room.CreatedRole,
		&room.UpdatedBy, &room.UpdatedName, &room.UpdatedEmail, &room.UpdatedRole,
		&room.CreatedAt, &room.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &room, nil
}

// GetAll retrieves all rooms
func (r *pgRoomRepository) GetAll() ([]domain.Room, error) {
	query := `SELECT 
		id, name, is_active,
		created_by, created_name, created_email, created_role,
		updated_by, updated_name, updated_email, updated_role,
		created_at, updated_at
	FROM rooms ORDER BY name`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var room domain.Room
		err := rows.Scan(
			&room.ID, &room.Name, &room.IsActive,
			&room.CreatedBy, &room.CreatedName, &room.CreatedEmail, &room.CreatedRole,
			&room.UpdatedBy, &room.UpdatedName, &room.UpdatedEmail, &room.UpdatedRole,
			&room.CreatedAt, &room.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
