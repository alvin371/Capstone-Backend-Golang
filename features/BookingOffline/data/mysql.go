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
	var record []OfflineClassUser
	err = book.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return ToBookingOfflineCoreList(record), nil
}
func (book *mysqlBookingOfflineClassRepo) InsertMemberBookingOffline(userID int, classID int) (err error) {
	var bookingoffline = OfflineClassUser{}
	bookingoffline.UserID = userID
	bookingoffline.ClassID = classID
	// convData := ToBookingOfflineCore(bookingoffline)

	if err := book.Conn.Create(&bookingoffline).Error; err != nil {
		return err
	}
	return nil
}
