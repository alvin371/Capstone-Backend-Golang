package rep

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
	Class     OnlineClassResponse
	User      UserResponse
}

type OnlineClassResponse struct {
	ID      int
	Name    string
	Day     string
	Date    string
	Link    string
	Time    string
	Trainer string
	Image   string
}
type UserResponse struct {
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

func ToUserResponse(data bookingOnline.User) UserResponse {
	return UserResponse{
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

func ToOfflineClassResponse(core bookingOnline.OnlineClassCore) OnlineClassResponse {
	return OnlineClassResponse{
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
func ToBookingOnlineCore(core bookingOnline.OnlineClassUser) OnlineClassUser {
	return OnlineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassResponse(core.Class),
		User:      ToUserResponse(core.User),
	}
}

func TobookingOnlineCoreList(oc []bookingOnline.OnlineClassUser) []OnlineClassUser {
	conv := []OnlineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, ToBookingOnlineCore(ocList))
	}
	return conv
}
