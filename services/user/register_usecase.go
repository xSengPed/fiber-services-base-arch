package userservice

import (
	"go-fiber-basic/datamodels"
	userrepository "go-fiber-basic/repository/user"
)

func RegisterUser(user *datamodels.User) error {
	err := userrepository.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}
