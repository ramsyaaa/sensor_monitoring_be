package models

import "time"

func (TokenAuth) TableName() string {
	return "token_auths"
}

type TokenAuth struct {
	AccessToken    string    `json:"access_token"`
	TokenType      string    `json:"token_type"`
	RefreshToken   string    `json:"refresh_token"`
	ExpiresIn      int64     `json:"expires_in"`
	Scope          string    `json:"scope"`
	ClientId       string    `json:"client_id"`
	ClientSecret   string    `json:"client_secret"`
	UserId         int64     `json:"user_id"`
	UpdateAt       time.Time `json:"update_at"`
	ExpiresAt      time.Time `json:"expires_at"`
	IsNewDev       bool      `json:"is_new_dev"`
	NewDevice      int64     `json:"new_device"`
	InsNewDevAt    time.Time `json:"insnewdev_at"`
	IsNewSensor    bool      `json:"is_new_sensor"`
	NewSensor      int64     `json:"new_sensor"`
	InsNewSensorAt time.Time `json:"insnewsensor_at"`
}
