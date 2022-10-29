package config

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
)

var mysqlConfig *DBConfig

func GetMysqlConfig() *DBConfig {
	if mysqlConfig == nil {
		mysqlConfig = initMysqlConfig()
	}

	return mysqlConfig
}

func initMysqlConfig() (newConfig *DBConfig) {
	/*
		initMysqlConfig: to initialize MySQL DB config for connection
	*/
	newConfig = &DBConfig{
		Address:  os.Getenv("MYSQL_HOST"),
		DBName:   os.Getenv("MYSQL_DBNAME"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}

	portEnv := os.Getenv("MYSQL_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatalf("cannot parse db port (must be valid int), err %v", err)
		return nil
	}

	newConfig.Port = port
	return
}
