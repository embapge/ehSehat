package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	Secret string
}

func NewManager(secret string) *Manager {
	return &Manager{Secret: secret}
}

func (j *Manager) GenerateToken(userID, name, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"name":  name,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

func (j *Manager) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
