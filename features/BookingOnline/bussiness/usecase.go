package bussiness

import (
	"capstone/backend/features/bookingOnline"
)

type bookingOnlineUseCase struct {
	bData bookingOnline.Data
}

func NewBussinessBookingOnlineClass(bOCData bookingOnline.Data) bookingOnline.Bussiness {
	return &bookingOnlineUseCase{
		bData: bOCData,
	}
}

func (booku *bookingOnlineUseCase) GetListBookingOnline(data bookingOnline.OnlineClassUser) (list []bookingOnline.OnlineClassUser, err error) {
	offlineClass, err := booku.bData.SelectAllBookingOnline(data)
	if err != nil {
		return nil, err
	}
	return offlineClass, nil
}

func (booku *bookingOnlineUseCase) MemberBookingOnline(userID int, classID int) (err error) {
	if err := booku.bData.InsertMemberBookingOnline(userID, classID); err != nil {
		return err
	}
	return nil
}
