package data

import (
	user "capstone/backend/features/User"
	"errors"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) user.Data {
	return &mysqlUserRepository{DB}
}

func (mr *mysqlUserRepository) InsertUserData(data user.Core) (int, error) {
	// recordData := toUserRecord(data)
	userData, _, _ := SeparateUserData(data)
	err := mr.DB.Create(&userData).Error

	if err != nil {
		return 0, err
	}

	return int(userData.ID), nil
}

func (mr *mysqlUserRepository) CheckUser(data user.Core) (user.Core, error) {
	var userData User
	err := mr.DB.Where("email = ? and password = ?", data.Email, data.Password).First(&userData).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.Core{}, errors.New("no existing user")
	}
	if err != nil {
		return user.Core{}, err
	}

	return toCore(userData), nil
}

func (mr *mysqlUserRepository) GetDataById(id int) (user.Core, error) {
	var userData User
	err := mr.DB.Preload("Skillsets").Preload("Experiences").First(&userData, id).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.Core{}, errors.New("no existing user")
	}
	if err != nil {
		return user.Core{}, err
	}

	return toCore(userData), nil
}

func (mr *mysqlUserRepository) UpdateUser(data user.Core) error {
	err := mr.DB.Debug().Model(&User{}).Where("id = ?", data.Id).Updates(User{
		Name:    data.Name,
		Dob:     data.Dob,
		Gender:  data.Gender,
		Address: data.Address,
		Title:   data.Title,
		Bio:     data.Bio,
		Email:   data.Email,
	}).Error
	if err != nil {
		return nil
	}

	return nil
}

func (mr *mysqlUserRepository) GetUserByEmail(email string) (bool, error) {
	var userModel User
	err := mr.DB.Where("email = ?", email).Find(&userModel).Error
	if err != nil {
		return false, err
	}
	if userModel.ID != 0 {
		return true, nil
	}
	return false, nil
}
