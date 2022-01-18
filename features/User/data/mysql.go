package data

import (
	user "capstone/backend/features/User"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserData struct {
	DB *gorm.DB
}

func NewMySqlUSer(DB *gorm.DB) user.Data {
	return &UserData{DB}
}

func (ud *UserData) InsertUser(data user.User) (usr user.User, err error) {
	record := fromCore(data)

	if err := ud.DB.Create(&record).Error; err != nil {
		return user.User{}, err
	}

	return data, nil
}

func (ud *UserData) SelectAllUser(search string) []user.User {
	var users []User

	err := ud.DB.Find(&users).Error

	if err != nil {
		return nil
	}
	return toUserCoreList(users)
}

func (ud *UserData) CheckAccount(data user.User) (user.User, error) {
	var users User
	err := ud.DB.Where("username = ? and password = ?", data.Username, data.Password).First(&users).Error

	// Eliminate null data
	if users.Username == "" && users.ID == 0 {
		return user.User{}, errors.New("user not found")
	}

	// Validate with DB
	if err != nil {
		return user.User{}, err
	}

	return toUserCore(users), nil
}

func (ud *UserData) UpdateUser(id int) (user.User, error) {
	var users User
	fmt.Println("Isi single Users : ", users)
	fmt.Println("id : ", id)
	err := ud.DB.Model(&users).Where("id=?", id).Updates(&users).Error
	// if users.Username == "" {
	// 	return user.User{}, errors.New("user not found")
	// }

	if err != nil {
		return user.User{}, err
	}

	return toUserCore(users), nil
}
