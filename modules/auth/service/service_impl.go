package service

import (
	"context"

	"sensor_monitoring_be/modules/auth/repository"
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
