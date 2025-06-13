package test

import (
	"backend/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestSentChat(t *testing.T) {
	DB, err := gorm.Open(mysql.Open("root:root123489+@tcp(127.0.0.1:3306)/raftsim?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}
	err = DB.AutoMigrate(
		&models.UserInformation{},
		&models.ChatInformation{},
	)
}
