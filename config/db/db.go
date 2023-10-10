package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/esb_legend?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("Database connection is nil")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database connection")
	}

	if err := sqlDB.Ping(); err != nil {
		panic("Database connection is closed")
	}

	return db
}
