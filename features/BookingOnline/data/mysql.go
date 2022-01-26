package data

import (
	"capstone/backend/features/bookingOnline"

	"gorm.io/gorm"
)

type mysqlbookingOnlineClassRepo struct {
	Conn *gorm.DB
}

func NewBookingOnlineRepository(conn *gorm.DB) bookingOnline.Data {
	return &mysqlbookingOnlineClassRepo{
		Conn: conn,
	}
}

func (book *mysqlbookingOnlineClassRepo) SelectAllBookingOnline(bookingOnline.OnlineClassUser) (list []bookingOnline.OnlineClassUser, err error) {
	var record []OnlineClassUser
	err = book.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return TobookingOnlineCoreList(record), nil
}
func (book *mysqlbookingOnlineClassRepo) InsertMemberBookingOnline(userID int, classID int) (err error) {
	var bookingOnline = OnlineClassUser{}
	bookingOnline.UserID = userID
	bookingOnline.ClassID = classID
	// convData := TobookingOnlineCore(bookingOnline)

	if err := book.Conn.Create(&bookingOnline).Error; err != nil {
		return err
	}
	return nil
}
