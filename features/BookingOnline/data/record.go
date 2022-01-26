package data

import (
	"capstone/backend/features/bookingOnline"
	"time"
)

type OnlineClassUser struct {
	ID        int
	ClassID   int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OnlineClassCore
	User      User
}

type OnlineClassCore struct {
	ID      int
	Name    string
	Day     string
	Date    string
	Link    string
	Time    string
	Trainer string
	Image   string
}
type User struct {
	ID           uint
	Username     string
	Role         string
	Email        string
	Password     string
	Token        string
	Avatar       string
	Goals        string
	MemberStatus string
}

func ToUserCore(data User) bookingOnline.User {
	return bookingOnline.User{
		ID:           data.ID,
		Username:     data.Username,
		Role:         data.Role,
		Email:        data.Email,
		Password:     data.Password,
		Token:        data.Token,
		Avatar:       data.Avatar,
		Goals:        data.Goals,
		MemberStatus: data.MemberStatus,
	}
}

func ToOnlineClassCore(core OnlineClassCore) bookingOnline.OnlineClassCore {
	return bookingOnline.OnlineClassCore{
		ID:      core.ID,
		Name:    core.Name,
		Day:     core.Day,
		Date:    core.Date,
		Link:    core.Link,
		Time:    core.Time,
		Trainer: core.Trainer,
		Image:   core.Image,
	}
}

func TobookingOnlineCoe(core OnlineClassUser) bookingOnline.OnlineClassUser {
	return bookingOnline.OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOnlineClassCore(core.Class),
		User:      ToUserCore(core.User),
	}
}

func TobookingOnlineCoreList(oc []OnlineClassUser) []bookingOnline.OnlineClassUser {
	conv := []bookingOnline.OnlineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, TobookingOnlineCoe(ocList))
	}
	return conv
}
