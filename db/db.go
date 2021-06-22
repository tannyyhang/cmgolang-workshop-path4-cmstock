package db

import (
	"main/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//ประกาศตัวแปร db แบบ private
var db *gorm.DB

// GetDB - call this method to get db
// GetDB เป็นตัวใหญ่เพราะว่าต้องการให้โลกภายนอกได้่ใช้
func GetDB() *gorm.DB{
	return db
}

// SetupDB - setup database for sharing to all api
func SetupDB() {
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.User{})
	db = database
}