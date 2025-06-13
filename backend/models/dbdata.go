package models

import (
	"gorm.io/gorm"
)

type ChatInformation struct {
	gorm.Model
	Username    string `gorm:"type:varchar(255);not null;uniqueIndex"`
	UserMessage string `gorm:"not null"`
	LLMMessage  string `gorm:"not null"`
}

type UserInformation struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Email    string `gorm:"type:varchar(255);not null"`
}

type UserLogin struct {
	Username string `json:"username"` // Username string
	Password string `json:"password"` // Password string
}

type UserRegister struct {
	Username string `json:"username"` // Username string
	Password string `json:"password"` // Password string
	Email    string `json:"email"`    // Email string
}
