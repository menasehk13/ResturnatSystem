package service

import (
	"fmt"
	"log"

	"github.com/menasehk13/ResturnatSystem/backend/config"
	"github.com/menasehk13/ResturnatSystem/backend/model"
)

func SaveUser(user *model.User) error {
	err := config.DB.Create(user).Error
	if err != nil {
		log.Printf("failed to save user to the database: %v", err.Error())
		return fmt.Errorf("failed to save user to the database: %v", err.Error())
	}

	return nil
}
