package registry

import (
	"log"

	"gorm.io/gorm"
)

var registeredModels []interface{}

func RegisterModel(model interface{}) {
    registeredModels = append(registeredModels, model)
}

func LoadModels(db *gorm.DB) {
	for _, model := range registeredModels {
        if err := migrate(db, model); err != nil {
            log.Fatal("failed to automigrate models:", err)
        }
    }
	log.Println("All models initialized successfully")
}

func migrate(db *gorm.DB, dst interface{}) error {
	// Check if the schema exists, and create it if it doesn't
	if !db.Migrator().HasTable(dst) {
		err := db.Migrator().CreateTable(dst)
		if err != nil {
			log.Println("Failed to create schema:", err)
			return err
		}
		log.Println("Schema created successfully")
	} else {
		log.Println("Schema already exists")
	}
	return nil
}

func init() {
	registeredModels = make([]interface{}, 0)
}
