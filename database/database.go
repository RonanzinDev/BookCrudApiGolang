package database

import (
	"log"
	"time"

	"github.com/ronanzindev/book-api-golang/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Iniciando o db
func StartDB() {
	str := "POSTGRES STRING CONNECTION HERE"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error :", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(10)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigration(db)

}

func GetDataBase() *gorm.DB {
	return db
}
