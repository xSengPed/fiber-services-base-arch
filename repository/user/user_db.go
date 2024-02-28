package userrepository

import (
	"go-fiber-basic/datamodels"
	"go-fiber-basic/repository"
)

func RegisterUser(user *datamodels.User) error {
	err := repository.DB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
