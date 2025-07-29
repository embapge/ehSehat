package domain

type AuthService interface {
	Register(name, email, password, role string) (*User, error)
	Login(email, password string) (string, error) // JWT token
}
