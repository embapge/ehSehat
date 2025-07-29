package infra

import (
	"context"
	"database/sql"
	"ehSehat/auth-service/internal/auth/domain"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type mySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *mySQLUserRepository {
	return &mySQLUserRepository{db}
}

func (r *mySQLUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {

	var users []*domain.User

	result, err := r.db.QueryContext(ctx, "SELECT id, name, email, role, password, created_at, updated_at FROM users")
	if err != nil {
		return []*domain.User{}, err
	}

	defer result.Close()

	for result.Next() {
		user := new(domain.User)
		result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		users = append(users, user)
	}

	return users, nil
}

func (r *mySQLUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, role, password, created_at, updated_at FROM users WHERE id = ? LIMIT 1", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

func (r *mySQLUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, role, password, created_at, updated_at FROM users WHERE email = ? LIMIT 1", email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *mySQLUserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	uuid := uuid.NewString()
	if uuid == "" {
		return &domain.User{}, errors.New("failed to generate UUID")
	}

	result, err := r.db.ExecContext(ctx,
		"INSERT INTO users (id, name, email, role, password) VALUES (?, ?, ?, ?, ?)",
		uuid, user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("no user inserted")
	}

	user.ID = uuid

	return user, nil
}

func (r *mySQLUserRepository) Update(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	result, _ := r.db.ExecContext(ctx, "UPDATE users SET name = ?, email = ?, role = ?, password = ? WHERE id = ?", user.Name, user.Email, user.Role, user.Password, user.ID)
	userID, _ := result.LastInsertId()

	if userID == 0 {
		return &domain.User{}, nil
	}

	user.ID = strconv.FormatInt(userID, 10)

	return user, nil
}

func (r *mySQLUserRepository) Delete(ctx context.Context, id string) error {
	result, _ := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("failed to delete user with id %v", id)
	}

	return nil
}
