package domain

import "time"

type User struct {
	ID        string // Adjust: ID jadi string (biar domain agnostic, conversion di infra)
	Name      string
	Email     string
	Role      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// --- Baru di DDD: Logic bisnis domain-level langsung di entity --- //
func (u *User) IsEmailValid() bool {
	return len(u.Email) > 5 // Sederhana, bisa ditingkatkan regex
}
