package service

import (
	"context"

	"sensor_monitoring_be/modules/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) Authenticate(ctx context.Context, username, password string) ([]map[string]interface{}, error) {
	return s.repo.Authenticate(ctx, username, password)
}

func (s *service) CreateUser(ctx context.Context, user map[string]interface{}) (map[string]interface{}, error) {
	// Hash the password using bcrypt
	password, ok := user["password"].(string)
	if ok {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user["password"] = string(hashedPassword)
	}

	// Set the role to visitor by default
	user["role"] = "visitor"

	return s.repo.CreateUser(ctx, user)
}

func (s *service) EditUser(ctx context.Context, id int, user map[string]interface{}) (map[string]interface{}, error) {
	// Hash the password using bcrypt
	password, ok := user["password"].(string)
	if ok {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user["password"] = string(hashedPassword)
	}

	return s.repo.EditUser(ctx, id, user)
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *service) CheckUsernameExist(ctx context.Context, username string) (bool, error) {
	return s.repo.CheckUsernameExist(ctx, username)
}

func (s *service) CheckPassword(ctx context.Context, username, password string) (bool, error) {
	return s.repo.CheckPassword(ctx, username, password)
}
func (s *service) ListUsers(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.ListUsers(ctx)
}
