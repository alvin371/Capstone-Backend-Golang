package data

import (
	"capstone/backend/features/bookingOffline"

	"gorm.io/gorm"
)

type mysqlBookingOfflineClassRepo struct {
	Conn *gorm.DB
}

func NewBookingOfflineRepository(conn *gorm.DB) bookingOffline.Data {
	return &mysqlBookingOfflineClassRepo{
		Conn: conn,
	}
}

func (book *mysqlBookingOfflineClassRepo) SelectAllBookingOffline(bookingOffline.OfflineClassUser) (list []bookingOffline.OfflineClassUser, err error) {

}
func (book *mysqlBookingOfflineClassRepo) InsertMemberBookingOffline(userID int, classID int) (err error) {

}
