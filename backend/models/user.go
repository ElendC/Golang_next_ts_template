package models

import "time"

type StudentUser struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"` //Don't need if we use FEIDE

	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
