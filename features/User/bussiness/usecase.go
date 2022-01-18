package bussiness

import (
	user "capstone/backend/features/User"
	"capstone/backend/middlewares"

	"github.com/go-playground/validator/v10"
)

type UserBussiness struct {
	userData user.Data
	validate *validator.Validate
}

func NewUserBussiness(userData user.Data) user.Bussiness {
	return &UserBussiness{
		userData: userData,
	}
}

func (ub *UserBussiness) GetAllUser(search string) (user []user.User) {
	user = ub.userData.SelectAllUser(search)
	return
}
func (ub *UserBussiness) CreateUser(data user.User) (resp user.User, err error) {
	if err := ub.validate.Struct(data); err != nil {
		return user.User{}, err
	}

	resp, err = ub.userData.InsertUser(data)
	if err != nil {
		return user.User{}, err
	}
	return user.User{}, nil
}
func (ub *UserBussiness) EditUser(id int) (usr user.User, err error) {
	userData, err := ub.userData.UpdateUser(id)

	if err != nil {
		return user.User{}, err
	}
	return userData, nil
}
func (ub *UserBussiness) LoginUser(user user.User) (usr user.User, err error) {
	accountData, err := ub.userData.CheckAccount(user)
	if err != nil {
		return user, err
	}
	accountData.Token, err = middlewares.CreateToken(user.ID, user.Username)
	if err != nil {
		return user, err
	}
	return accountData, nil
}
