package bookingOffline

import "time"

type OfflineClassUser struct {
	ID        int
	ClassID   int `gorm:"primaryKey"`
	UserID    int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassCore
	User      User
}

type OfflineClassCore struct {
	ID        int
	Name      string `json:"name"`
	Day       string `json:"day"`
	Date      string `json:"date"`
	Location  string `json:"location"`
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
	GetListBookingOffline(OfflineClassUser) (list []OfflineClassUser, err error)
	MemberBookingOffline(userID int, classID int) (err error)
}

type Data interface {
	SelectAllBookingOffline(OfflineClassUser) (list []OfflineClassUser, err error)
	InsertMemberBookingOffline(userID int, classID int) (err error)
}
