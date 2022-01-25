package rep

import (
	user "capstone/backend/features/User"
)

type User struct {
	Username string `json: "username"`
	Role     string `json: "role"`
	Email    string `json: "email"`
	Password string `json: "password"`
	Token    string `json:"token"`
}

type UserLogin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
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
