package data

import (
	user "capstone/backend/features/User"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserData struct {
	DB *gorm.DB
}

func NewMySqlUSer(DB *gorm.DB) user.Data {
	return &UserData{DB}
}

func (ud *UserData) InsertUser(data user.User) (err error) {
	record := fromCore(data)
	if err := ud.DB.Create(&record).Error; err != nil {
		return err
	}

	return nil
}
func (ud *UserData) SelectUserById(id int) (user.User, error) {
	var userID User

	err := ud.DB.Find(&userID, id).Error

	if userID.Username == "" {
		return user.User{}, err
	}

	if err != nil {
		return user.User{}, err
	}
	return toUserCore(userID), err
}

func (ud *UserData) SelectAllUser(userData user.User) ([]user.User, error) {
	var users []User

	err := ud.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}
	return toUserCoreList(users), err
}

func (ud *UserData) CheckAccount(data user.User) (user.User, error) {
	var userAuth User
	err := ud.DB.Where("username = ?", data.Username).First(&userAuth).Error
	// Eliminate null data
	if userAuth.Username == "" && userAuth.ID == 0 {
		return user.User{}, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(data.Password)); err != nil {
		log.Println(err)
		return user.User{}, errors.New("Password doesnt match")
	}
	fmt.Println(err)

	// Validate with DB
	if err != nil {
		return user.User{}, err
	}

	return toUserCore(userAuth), nil
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
