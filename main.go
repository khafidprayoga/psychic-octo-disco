package main

import (
	"github.com/khafidprayoga/psychic-octo-disco/config"
	"github.com/khafidprayoga/psychic-octo-disco/database"
	"github.com/khafidprayoga/psychic-octo-disco/registry"
)

// Mount all application backend dependencies
func main() {
	// Mysql Config
	mysqlCfg := config.GetMysqlConfig()
	mysqlClient := database.InitMySqlDB(mysqlCfg)

	registry.StartBackend(config.Server, *mysqlClient)
}
