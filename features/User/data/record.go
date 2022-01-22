package data

import (
	user "capstone/backend/features/User"

	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Username     string    `json: "username" gorm:"unique;not null" form: "username"`
	Password     string    `json: "password" gorm:"not null" form: "password"`
	Role         string    `json: "role" form:"role"`
	Email        string    `json: "email" gorm:"unique;not null" form: "email"`
	Avatar       string    `json: "avatar" form: "avatar"`
	Goals        string    `json:"goals" form: "goals"`
	MemberStatus string    `json: "member_status" form: "member_status"`
}

func toUserRecord(user user.User) User {
	return User{

		ID: user.ID,

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
