package main

import (
	"fmt"
	"github.com/khafidprayoga/psychic-octo-disco/config"
	"github.com/khafidprayoga/psychic-octo-disco/database"
	"github.com/khafidprayoga/psychic-octo-disco/database/tables"
	"gorm.io/gorm"
	"log"
)

// Play on localhost instead on go.dev/play
func main() {
	//fmt.Println("Hello World")
	mysqlCfg := config.GetMysqlConfig()
	mysqlClient := database.InitMySqlDB(mysqlCfg)
	database.MigrateUp(mysqlClient)
	//database.MigrateDown(mysqlClient)
	Seed(mysqlClient)
}

func Seed(db *gorm.DB) {

	log.Println("seeding dummy data")
	priority := []string{"high", "very-high"}
	for i := 100; i < 150; i++ {
		active := i % 2
		todo := tables.Todo{
			Title:           fmt.Sprintf("Todo hari ke-%v", i),
			ActivityGroupId: i%2 + 1,
			Priority:        &priority[i%2],
			IsActive:        &active,
		}

		db.Create(&todo)
	}
	email := "root@localhost"
	db.Create(&tables.Activity{
		ID:    1,
		Title: "aktivitas kantor",
		Email: &email,
	})
	db.Create(&tables.Activity{
		ID:    2,
		Title: "aktivitas rumah",
		Email: &email,
	})
	log.Println("finished insert dummy data")
}
