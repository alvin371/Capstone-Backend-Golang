package data

import (
	"capstone/backend/features/bookingOffline"
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   OfflineClassCore
	UserID    User
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassCore
	User      User
}

type OfflineClassCore struct {
	ID       int
	Name     string
	Day      string
	Date     string
	Location string
	Time     string
	Trainer  string
	Image    string
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

func ToUserCore(data User) bookingOffline.User {
	return bookingOffline.User{
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

func ToOfflineClassCore(core OfflineClassCore) bookingOffline.OfflineClassCore {
	return bookingOffline.OfflineClassCore{
		ID:       core.ID,
		Name:     core.Name,
		Day:      core.Day,
		Date:     core.Date,
		Location: core.Location,
		Time:     core.Time,
		Trainer:  core.Trainer,
		Image:    core.Image,
	}
}

func ToBookingOfflineCore(core OfflineClassUser) bookingOffline.OfflineClassUser {
	return bookingOffline.OfflineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassCore(core.Class),
		User:      ToUserCore(core.User),
	}
}
