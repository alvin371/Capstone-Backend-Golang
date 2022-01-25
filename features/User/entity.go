package user

import "time"

type User struct {
	ID           uint
	Username     string `json:"username" gorm:"unique;not null"`
	Role         string `json:"role" gorm:"default:user"`     //user
	Email        string `json:"email" gorm:"unique;not null"` //unique,
	Password     string `json:"password" gorm:"not null"`
	Token        string `jason:"token" form:"token"`
	Avatar       string `json:"avatar" form:"avatar"`
	Goals        string `json:"goals" form:"goals"`
	MemberStatus string `json:"member_status" gorm:"default:reguler"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Bussiness interface {
	GetAllUser(User) (user []User, err error)
	GetUserById(id int) (user User, err error)
	CreateUser(data User) (err error)
	EditUser(id int) (usr User, err error)
	LoginUser(user User) (usr User, err error)
}

type Data interface {
	SelectAllUser(User) (user []User, err error)
	SelectUserById(id int) (user User, err error)
	InsertUser(data User) (err error)
	UpdateUser(id int) (usr User, err error)
	CheckAccount(User) (user User, err error)
}
