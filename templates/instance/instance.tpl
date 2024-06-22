package instance

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

	"{{.PackageName}}/config"
)

type Instance interface {
	Destroy() error
    DB() *gorm.DB
}

type instance struct {
    db *gorm.DB
	// TODO: define singleton object here
}

// Destroy closes the connections & cleans up the instance
func (instance *instance) Destroy() error {
	return nil
}

func (instance *instance) DB() *gorm.DB {
	return instance.db
}

// Init initializes the instance
func Init(cfg config.Configuration) Instance {
	instance := &instance{}
    
    gormDB, err := gorm.Open(sqlite.Open(cfg.DatabaseName()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	instance.db = gormDB

    // TODO: initialize singleton object here

	return instance
}
