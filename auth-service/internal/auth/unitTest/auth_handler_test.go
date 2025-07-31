package unitTest

import (
	"context"
	"errors"
	"testing"

	grpcHandler "ehSehat/auth-service/internal/auth/delivery/grpc"
	"ehSehat/auth-service/internal/auth/delivery/grpc/pb"
	"ehSehat/auth-service/internal/auth/domain"

	"ehSehat/auth-service/internal/auth/app"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// -----------------------------
// ✅ Mock: UserRepository
// -----------------------------
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if user, ok := args.Get(0).(*domain.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if created, ok := args.Get(0).(*domain.User); ok {
		return created, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepo) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Add missing GetAll method to satisfy domain.UserRepository interface
func (m *MockUserRepo) GetAll(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	if users, ok := args.Get(0).([]*domain.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

// Add missing GetByID method to satisfy domain.UserRepository interface
func (m *MockUserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*domain.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// Add missing Update method to satisfy domain.UserRepository interface
func (m *MockUserRepo) Update(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, id, user)
	if updated, ok := args.Get(0).(*domain.User); ok {
		return updated, args.Error(1)
	}
	return nil, args.Error(1)
}

// -----------------------------
// ✅ Mock: PasswordHasher
// -----------------------------
type MockHasher struct {
	mock.Mock
}

func (m *MockHasher) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockHasher) Check(raw string, hashed string) bool {
	args := m.Called(raw, hashed)
	return args.Bool(0)
}

// -----------------------------
// ✅ Mock: JWTManager
// -----------------------------
type MockJWTManager struct {
	mock.Mock
}

func (m *MockJWTManager) GenerateToken(id, name, email, role string) (string, error) {
	args := m.Called(id, name, email, role)
	return args.String(0), args.Error(1)
}

// -----------------------------
// ✅ Tests (No handler change)
// -----------------------------
func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHasher := new(MockHasher)
	mockJWT := new(MockJWTManager)

	appInstance := app.NewAuthApp(mockRepo, mockHasher, mockJWT, &amqp.Channel{}) // 🧩 pakai real struct
	handler := grpcHandler.NewAuthHandler(appInstance)

	req := &pb.RegisterRequest{
		Name:     "John",
		Email:    "john@example.com",
		Password: "secret",
		Role:     "patient",
	}

	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(nil, nil)
	mockHasher.On("Hash", req.Password).Return("hashed-secret", nil)

	createdUser := &domain.User{ID: "1", Name: req.Name, Email: req.Email, Role: req.Role}
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(createdUser, nil)

	res, err := handler.Register(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "1", res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Email, res.Email)
	assert.Equal(t, req.Role, res.Role)
}

func TestRegister_EmailExist(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHasher := new(MockHasher)
	mockJWT := new(MockJWTManager)

	appInstance := app.NewAuthApp(mockRepo, mockHasher, mockJWT, &amqp.Channel{})
	handler := grpcHandler.NewAuthHandler(appInstance)

	req := &pb.RegisterRequest{
		Name:     "Jane",
		Email:    "jane@example.com",
		Password: "pass",
		Role:     "doctor",
	}

	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(&domain.User{ID: "existing"}, nil)

	res, err := handler.Register(context.Background(), req)
	assert.Nil(t, res)
	assert.Equal(t, codes.AlreadyExists, status.Code(err))
	assert.Contains(t, err.Error(), "email sudah terdaftar")
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHasher := new(MockHasher)
	mockJWT := new(MockJWTManager)

	appInstance := app.NewAuthApp(mockRepo, mockHasher, mockJWT, &amqp.Channel{})
	handler := grpcHandler.NewAuthHandler(appInstance)

	req := &pb.LoginRequest{
		Email:    "john@example.com",
		Password: "secret",
	}

	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(&domain.User{
		ID:       "1",
		Name:     "John",
		Email:    req.Email,
		Role:     "patient",
		Password: "hashed-secret",
	}, nil)

	mockHasher.On("Check", req.Password, "hashed-secret").Return(true)
	mockJWT.On("GenerateToken", "1", "John", req.Email, "patient").Return("mocked-token", nil)

	res, err := handler.Login(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Email, res.Email)
	assert.Equal(t, "mocked-token", res.Token)
}

func TestLogin_Unauthorized_WrongPassword(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHasher := new(MockHasher)
	mockJWT := new(MockJWTManager)

	appInstance := app.NewAuthApp(mockRepo, mockHasher, mockJWT, &amqp.Channel{})
	handler := grpcHandler.NewAuthHandler(appInstance)

	req := &pb.LoginRequest{
		Email:    "john@example.com",
		Password: "wrongpassword",
	}

	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(&domain.User{
		ID:       "1",
		Name:     "John",
		Email:    req.Email,
		Role:     "patient",
		Password: "hashed-secret",
	}, nil)

	mockHasher.On("Check", req.Password, "hashed-secret").Return(false)

	res, err := handler.Login(context.Background(), req)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unauthenticated, status.Code(err))
	assert.Contains(t, err.Error(), "email/password salah")
}

func TestLogin_Failure_DBError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockHasher := new(MockHasher)
	mockJWT := new(MockJWTManager)

	appInstance := app.NewAuthApp(mockRepo, mockHasher, mockJWT, &amqp.Channel{})
	handler := grpcHandler.NewAuthHandler(appInstance)

	req := &pb.LoginRequest{
		Email:    "error@example.com",
		Password: "pass",
	}

	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(nil, errors.New("db error"))

	res, err := handler.Login(context.Background(), req)
	assert.Nil(t, res)
	assert.Equal(t, codes.Unauthenticated, status.Code(err))
	assert.Contains(t, err.Error(), "email/password salah")
}
