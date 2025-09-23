/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package db

import (
	"log"

	"github.com/goretante/go-cli-to-do/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
		return
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
