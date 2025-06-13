package models

import (
	"backend/config"
	"fmt"
	"log"
)

func GetUserNameAndPwd(username string) (string, string) {
	var user UserInformation
	config.DB.First(&user, "username = ?", username)
	log.Println("Retrieved user:", user.Username, "with password:", user.Password)
	return user.Username, user.Password
}

func CreateUser(user UserInformation) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	log.Println("User created successfully:", user.Username)
	return nil
}

func GetUserByUsername(username string) (*UserInformation, error) {
	var user UserInformation
	if err := config.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}
