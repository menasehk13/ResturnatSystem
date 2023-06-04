package service

import (
	"fmt"

	"github.com/menasehk13/ResturnatSystem/backend/config"
	"github.com/menasehk13/ResturnatSystem/backend/model"
)

func SaveMenu(menu model.Menu) error {
	err := config.DB.Create(menu).Error

	if err != nil {
		return fmt.Errorf("error creating menu",err.Error())
	}
	return nil
}