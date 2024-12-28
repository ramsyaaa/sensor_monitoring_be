package service

import (
	"context"
)

type AuthService interface {
	Authenticate(ctx context.Context, username, password string) ([]map[string]interface{}, error)
	CreateUser(ctx context.Context, user map[string]interface{}) (map[string]interface{}, error)
	EditUser(ctx context.Context, id int, user map[string]interface{}) (map[string]interface{}, error)
	DeleteUser(ctx context.Context, id int) error
	CheckUsernameExist(ctx context.Context, username string) (bool, error)
	CheckPassword(ctx context.Context, username, password string) (bool, error)
}
