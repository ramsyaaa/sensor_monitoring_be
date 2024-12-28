package repository

import (
	"context"
	"errors"
	"sensor_monitoring_be/models"

	"golang.org/x/crypto/bcrypt"
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
	err := r.db.WithContext(ctx).Model(&models.TokenAuth{}).Where("id = ?", 1).Select("access_token, client_id, client_secret, expires_at, user_id").Find(&auth).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return auth, err
}

func (r *repository) CreateUser(ctx context.Context, user map[string]interface{}) (map[string]interface{}, error) {
	var newUser models.TokenAuth
	err := r.db.WithContext(ctx).Error
	if err != nil {
		return nil, err
	}

	// Assign values from the user map to the newUser struct
	newUser.Username = user["username"].(string)
	newUser.Password = user["password"].(string)
	newUser.Role = user["role"].(string)

	err = r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) EditUser(ctx context.Context, id int, user map[string]interface{}) (map[string]interface{}, error) {
	var existingUser models.TokenAuth
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	// Hanya update password jika ada dalam map user
	if password, ok := user["password"]; ok {
		err = r.db.WithContext(ctx).Model(&existingUser).Update("password", password).Error
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *repository) DeleteUser(ctx context.Context, id int) error {
	var user models.TokenAuth
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CheckUsernameExist(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.TokenAuth{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) CheckPassword(ctx context.Context, username, password string) (bool, error) {
	var user models.TokenAuth
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
func (r *repository) ListUsers(ctx context.Context) ([]map[string]interface{}, error) {
	var users []models.TokenAuth
	err := r.db.WithContext(ctx).Where("role = ?", "visitor").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var userList []map[string]interface{}
	for _, user := range users {
		userList = append(userList, map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
		})
	}

	return userList, nil
}
