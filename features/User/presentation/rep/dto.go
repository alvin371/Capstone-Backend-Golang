package rep

import (
	user "capstone/backend/features/User"
)

type User struct {
	Username     string `json:"username"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	Avatar       string `json:"avatar"`
	Goals        string `json:"goals"`
	MemberStatus string `json:"member_status"`
}

type UserLogin struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	Goals        string `json:"goals"`
	MemberStatus string `json:"member_status"`
}

func ToUserCore(req user.User) User {
	return User{
		Username: req.Username,
		Role:     req.Role,
		Email:    req.Email,
		Password: req.Password,
		Token:    req.Token,
	}
}

func ToUserCoreSlice(core []user.User) []User {
	var userArray []User
	for key := range core {
		userArray = append(userArray, ToUserCore(core[key]))
	}
	return userArray
}

func ToUserLoginResponse(user user.User) UserLogin {
	return UserLogin{
		ID:       user.ID,
		Username: user.Username,
		Token:    user.Token,
	}
}
