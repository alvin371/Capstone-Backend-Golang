package bookingOnline

import "time"

type OnlineClassUser struct {
	ID        int
	ClassID   int `gorm:"primaryKey"`
	UserID    int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OnlineClassCore
	User      User
}

type OnlineClassCore struct {
	ID        int
	Name      string `json:"name"`
	Day       string `json:"day"`
	Date      string `json:"date"`
	Link      string `json:"link"`
	Time      string `json:"time"`
	Trainer   string `json:"trainer"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
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
	GetListBookingOnline(OnlineClassUser) (list []OnlineClassUser, err error)
	MemberBookingOnline(userID int, classID int) (err error)
}

type Data interface {
	SelectAllBookingOnline(OnlineClassUser) (list []OnlineClassUser, err error)
	InsertMemberBookingOnline(userID int, classID int) (err error)
}
