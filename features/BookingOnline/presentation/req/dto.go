package req

import (
	"capstone/backend/features/bookingOnline"
)

type OnlineClassUser struct {
	ClassID int `json:"class_id"`
	UserID  int `json:"user_id"`
}

func (app *OnlineClassUser) ToCore() bookingOnline.OnlineClassUser {
	return bookingOnline.OnlineClassUser{
		ClassID: app.ClassID,
		UserID:  app.UserID,
	}
}
