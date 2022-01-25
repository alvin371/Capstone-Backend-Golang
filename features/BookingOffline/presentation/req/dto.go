package req

import "capstone/backend/features/bookingOffline"

type OfflineClassUser struct {
	ClassID int `json:"class_id"`
	UserID  int `json:"user_id"`
}

func (app *OfflineClassUser) ToCore() bookingOffline.OfflineClassUser {
	return bookingOffline.OfflineClassUser{
		ClassID: app.ClassID,
		UserID:  app.UserID,
	}
}
