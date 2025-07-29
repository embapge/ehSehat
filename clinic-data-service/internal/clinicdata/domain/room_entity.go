package domain

import "time"

// Room represents the room entity in the domain layer
type Room struct {
	ID       string
	Name     string
	IsActive bool

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

// RoomRepository defines repository behavior for room entity
type RoomRepository interface {
	Create(room *Room) (*Room, error)
	GetByID(id string) (*Room, error)
	GetAll() ([]Room, error)
}
