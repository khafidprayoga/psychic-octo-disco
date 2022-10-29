package database

import (
	"fmt"
	"github.com/khafidprayoga/psychic-octo-disco/config"
	"github.com/khafidprayoga/psychic-octo-disco/database/tables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMySqlDB(cfg *config.DBConfig) *gorm.DB {
	logLevel := logger.Silent
	if os.Getenv("DEBUG") == "1" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logLevel,        // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	sqlString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(sqlString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(
		&tables.Todo{},
		&tables.Activity{},
	)
}

func MigrateDown(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&tables.Todo{},
		&tables.Activity{},
	)
}
