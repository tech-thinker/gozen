package models

import (
	"log"

	"gorm.io/gorm"
)

func RegisterModels(db *gorm.DB) {
	// Automigrate the models
	if err := migrate(db, &Health{}); err != nil {
		log.Fatal("failed to automigrate models:", err)
	}
}

func migrate(db *gorm.DB, dst interface{}) error {
	// Check if the table exists, and create it if it doesn't
	if !db.Migrator().HasTable(dst) {
		err := db.Migrator().CreateTable(dst)
		if err != nil {
			log.Println("Failed to create table:", err)
			return err
		}
		log.Println("Table created successfully")
	} else {
		log.Println("Table already exists")
	}
	return nil
}
