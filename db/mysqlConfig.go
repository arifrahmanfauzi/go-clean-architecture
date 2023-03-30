package db

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	processENV()
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		defer logrus.Info("Cannot connect into database")
		logrus.Error("Cannot connect to mysql" + err.Error())
	}

	return db
}

func processENV() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error load .env file")
	}
}
