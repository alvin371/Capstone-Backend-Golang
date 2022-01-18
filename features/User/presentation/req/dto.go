package req

import (
	user "capstone/backend/features/User"
)

type User struct {
	ID         uint
	Username   string `json: "username"`
	Role       string `json: "role"`
	Email      string `json: "email"`
	Password   string `json: "password"`
	Token      string `json:"token"`
	Created_at string `json: "created_at"`
	Updated_at string `json: "updated_at"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromCore(core User) user.User {
	return user.User{
		ID:         core.ID,
		Username:   core.Username,
		Role:       core.Role,
		Email:      core.Email,
		Password:   core.Password,
		Token:      core.Token,
		Created_at: core.Created_at,
		Updated_at: core.Updated_at,
	}
}

func (userA *UserAuth) ToUserAuth() user.User {
	return user.User{
		Username: userA.Username,
		Password: userA.Password,
	}
}
