package data

import (
	user "capstone/backend/features/User"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func toUserRecord(user user.User) User {
	return User{
		Model: gorm.Model{
			ID: user.ID,
		},
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}

func toUserCore(usr User) user.User {
	return user.User{
		ID:       usr.ID,
		Username: usr.Username,
		Password: usr.Password,
		Email:    usr.Email,
	}
}

func toUserCoreList(accList []User) []user.User {
	convUser := []user.User{}

	for _, user := range accList {
		convUser = append(convUser, toUserCore(user))
	}
	return convUser
}

func fromCore(usr user.User) User {
	return User{}
}
