package app

import (
	"context"
	"ehSehat/auth-service/internal/auth/domain"
)

type AuthApp struct {
	UserRepo       domain.UserRepository
	PasswordHasher PasswordHasher
	JWTManager     JWTManager
}

func NewAuthApp(repo domain.UserRepository, hasher PasswordHasher, jwt JWTManager) *AuthApp {
	return &AuthApp{
		UserRepo:       repo,
		PasswordHasher: hasher,
		JWTManager:     jwt,
	}
}

// Register user baru
func (a *AuthApp) Register(ctx context.Context, name, email, password, role string) (*domain.User, error) {
	if name == "" || email == "" || password == "" || role == "" {
		return nil, ErrValidation
	}
	existing, _ := a.UserRepo.GetByEmail(ctx, email)
	if existing != nil && existing.ID != "" {
		return nil, ErrEmailExist
	}
	hashed, err := a.PasswordHasher.Hash(password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Name:     name,
		Email:    email,
		Role:     role,
		Password: hashed,
	}

	createdUser, err := a.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

// Login user
func (a *AuthApp) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.UserRepo.GetByEmail(ctx, email)

	if err != nil || user == nil {
		return "", ErrUnauthorized
	}

	if !a.PasswordHasher.Check(password, user.Password) {
		return "", ErrUnauthorized
	}
	token, err := a.JWTManager.GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
