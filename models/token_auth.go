package models

import "time"

func (TokenAuth) TableName() string {
	return "token_auths"
}

type TokenAuth struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	Scope        string    `json:"scope"`
	ClientId     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	UserId       int64     `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}
