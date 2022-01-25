package req

import (
	user "capstone/backend/features/User"
	"time"
)

type User struct {
	ID           uint
	Username     string `form:"username"`
	Password     string `form:"password"`
	Role         string `form:"role"`
	Email        string `form:"email"`
	Token        string
	Avatar       string `form:"avatar"`
	Goals        string `form:"goals"`
	MemberStatus string `form:"member_status"`
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
