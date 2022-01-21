package req

import (
	user "capstone/backend/features/User"
	"time"
)

type User struct {
	ID           uint
	Username     string    `json: "username" form:"username"`
	Role         string    `json: "role" form:"role"`
	Email        string    `json: "email" form:"email"`
	Password     string    `json: "password" form:"password"`
	Token        string    `json:"token" form:"token"`
	Avatar       string    `json: "avatar" form:"avatar"`
	Goals        string    `json:"goals" form:"goals"`
	MemberStatus string    `json: "member_status" form:"member_status"`
	Created_at   time.Time `json: "created_at"`
	Updated_at   time.Time `json: "updated_at"`
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
		Created_at:   core.Created_at,
		Updated_at:   core.Updated_at,
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
		Created_at:   core.Created_at,
		Updated_at:   core.Updated_at,
	}
}

func (userA *UserAuth) ToUserAuth() user.User {
	return user.User{
		Username: userA.Username,
		Password: userA.Password,
	}
}
