package repository

import (
	"context"
)

type AuthRepository interface {
	Authenticate(ctx context.Context, username, password string) ([]map[string]interface{}, error)
}
