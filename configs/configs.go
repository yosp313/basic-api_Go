package configs

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConfig() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to db")

	return db
}
