package rep

import (
	"capstone/backend/features/bookingOffline"
	"time"
)

type OfflineClassUser struct {
	ID        int
	ClassID   int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Class     OfflineClassResponse
	User      UserResponse
}

type OfflineClassResponse struct {
	ID       int
	Name     string
	Day      string
	Date     string
	Location string
	Time     string
	Trainer  string
	Image    string
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

func ToUserResponse(data bookingOffline.User) UserResponse {
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

func ToOfflineClassResponse(core bookingOffline.OfflineClassCore) OfflineClassResponse {
	return OfflineClassResponse{
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
func ToBookingOfflineCore(core bookingOffline.OfflineClassUser) OfflineClassUser {
	return OfflineClassUser{
		ID:        core.ID,
		ClassID:   core.Class.ID,
		UserID:    int(core.User.ID),
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Class:     ToOfflineClassResponse(core.Class),
		User:      ToUserResponse(core.User),
	}
}

func ToBookingOfflineCoreList(oc []bookingOffline.OfflineClassUser) []OfflineClassUser {
	conv := []OfflineClassUser{}

	for _, ocList := range oc {
		conv = append(conv, ToBookingOfflineCore(ocList))
	}
	return conv
}
