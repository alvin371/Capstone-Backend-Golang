package data

import (
	user "capstone/backend/features/User"
	"time"
)

type User struct {
	ID           uint
	Username     string `json: "username" form: "username" gorm:"unique;not null"`
	Password     string `json: "password" form: "password" gorm:"not null"`
	Role         string `json: "role" form:"role" gorm:"default:user"`
	Email        string `json: "email" form: "email" gorm:"unique;not null"`
	Avatar       string `json: "avatar" form: "avatar"`
	Goals        string `json:"goals" form: "goals"`
	MemberStatus string `json: "member_status" form: "member_status" gorm:"default:reguler"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	return User{
		Username:     usr.Username,
		Email:        usr.Email,
		Role:         usr.Role,
		Password:     usr.Password,
		Avatar:       usr.Avatar,
		Goals:        usr.Goals,
		MemberStatus: usr.MemberStatus,
	}
}
