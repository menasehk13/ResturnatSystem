package config

import (
	"fmt"

	"github.com/menasehk13/ResturnatSystem/backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	// Connect to the database
	dsn := "root:@tcp(localhost:3306)/resturantsystem"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate the database: %v", err)
	}
	DB = db

	return nil
}

func CloseDatabase() error {
	db, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying DB connection: %v", err)
	}

	return db.Close()
}