package req

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
	Token        string `json:"token"`
	Avatar       string `json: "avatar" form: "avatar"`
	Goals        string `json:"goals" form: "goals"`
	MemberStatus string `json: "member_status" form: "member_status" gorm:"default:reguler"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromCore(core User) user.User {
	return user.User{
		ID:           core.ID,
		Username:     core.Username,
		Role:         core.Role,
		Email:        core.Email,
		Password:     core.Password,
		Token:        core.Token,
		Avatar:       core.Avatar,
		Goals:        core.Goals,
		MemberStatus: core.MemberStatus,
		Created_at:   core.CreatedAt,
		Updated_at:   core.UpdatedAt,
	}
}

func (core *User) ToUserCore() user.User {
	return user.User{
		ID:           core.ID,
		Username:     core.Username,
		Role:         core.Role,
		Email:        core.Email,
		Password:     core.Password,
		Token:        core.Token,
		Avatar:       core.Avatar,
		Goals:        core.Goals,
		MemberStatus: core.MemberStatus,
		Created_at:   core.CreatedAt,
		Updated_at:   core.UpdatedAt,
	}
}

func (userA *UserAuth) ToUserAuth() user.User {
	return user.User{
		Username: userA.Username,
		Password: userA.Password,
	}
}
