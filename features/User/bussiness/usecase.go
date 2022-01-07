package bussiness

import (
	user "capstone/backend/features/User"
	helper "capstone/backend/helpers"
	"errors"
	"fmt"
)

type userUseCase struct {
	userData user.Data
}

func NewUserBussiness(uData user.Data) user.Bussiness {
	return &user.Bussiness{userUseCase}
}

func (us *userUseCase) RegisterUser(data user.Core) error {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) || len(data.Name) == 0 {
		return errors.New("incomplete or invalid data")
	}

	isExist, err := us.userData.GetUserByEmail(data.Email)
	if err != nil {
		return err
	}
	if isExist {
		msg := fmt.Sprintf("email %v already in used", data.Email)
		return errors.New(msg)
	}

	userId, err := us.userData.InsertUserData(data)
	if err != nil {
		return err
	}

	return nil
}
