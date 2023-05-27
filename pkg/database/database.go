package database

import (
	"fmt"
	"simpleCRM/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConn() {
	var err error

	Db, err = gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
	if err != nil {
		panic("database connection failed")
	}
	fmt.Println("database conncted successfully")
	Db.AutoMigrate(&models.Person{})
	fmt.Println("database migrated")
}
