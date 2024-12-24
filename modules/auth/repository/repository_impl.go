package repository

import (
	"context"
	"errors"
	"sensor_monitoring_be/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{db: db}
}

func (r *repository) Authenticate(ctx context.Context, username, password string) ([]map[string]interface{}, error) {
	var auth []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&models.TokenAuth{}).Where("user_id = ?", 99837).Select("access_token, client_id, client_secret, expires_at, user_id").Find(&auth).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return auth, err
}
