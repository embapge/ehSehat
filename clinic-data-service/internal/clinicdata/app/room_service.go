package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"log"
	"strings"
)

// RoomService interface untuk business logic room
type RoomService interface {
	Create(room *domain.Room) (*domain.Room, error)
	GetByID(id string) (*domain.Room, error)
	GetAll() ([]domain.Room, error)
}

// roomService implements RoomService
type roomService struct {
	repo domain.RoomRepository
}

// NewRoomService constructor
func NewRoomService(r domain.RoomRepository) RoomService {
	return &roomService{repo: r}
}

// Create room dengan validasi nama wajib dan unik
func (s *roomService) Create(r *domain.Room) (*domain.Room, error) {
	log.Printf("DEBUG create room: Name='%s'", r.Name)

	if strings.TrimSpace(r.Name) == "" {
		return nil, ErrMissingFields
	}

	// Cek duplikasi nama (pakai GetAll karena tidak ada GetByName)
	rooms, err := s.repo.GetAll()
	if err != nil {
		log.Printf("ERROR get all rooms for duplicate check: %v", err)
		return nil, ErrInternal
	}
	for _, existing := range rooms {
		if strings.EqualFold(existing.Name, r.Name) {
			return nil, ErrInvalidInput // bisa diganti error custom "room already exists"
		}
	}

	return s.repo.Create(r)
}

// GetByID returns a room by ID
func (s *roomService) GetByID(id string) (*domain.Room, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrMissingID
	}
	return s.repo.GetByID(id)
}

// GetAll returns all rooms
func (s *roomService) GetAll() ([]domain.Room, error) {
	return s.repo.GetAll()
}
