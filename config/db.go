package config

import (
	"fmt"
	"gorillaz-clean-v3/helpers"
	"gorillaz-clean-v3/models"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func Init() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		helpers.Logger("error", "Error getting env")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	conn, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
	}

	conn.Debug().AutoMigrate(&models.User{})

	return conn
}
