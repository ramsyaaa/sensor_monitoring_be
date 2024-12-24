package service

import (
	"context"
)

type AuthService interface {
	Authenticate(ctx context.Context, username, password string) ([]map[string]interface{}, error)
}
