package instance

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

	"{{.PackageName}}/config"
	"{{.PackageName}}/constants"
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
    
    var conn gorm.Dialector
    if cfg.DbDriver() == constants.DbDriverSQLite {
        conn = sqlite.Open(cfg.DbName())
    } else if cfg.DbDriver() == constants.DbDriverMySQL {
        url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DbUser(), cfg.DbPass(), cfg.DbHost(), cfg.DbPort(), cfg.DbName())
        conn = mysql.Open(url)
    } else if cfg.DbDriver() == constants.DbDriverPostgres {
        url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.DbHost(), cfg.DbUser(), cfg.DbPass(), cfg.DbName(), cfg.DbPort())
        conn = postgres.Open(url)
    }
    gormDB, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	instance.db = gormDB

    // TODO: initialize singleton object here

	return instance
}
