package bussiness

import (
	"capstone/backend/features/bookingOffline"
)

type bookingOfflineUseCase struct {
	bData bookingOffline.Data
}

func NewBussinessBookingOfflineClass(bOCData bookingOffline.Data) bookingOffline.Bussiness {
	return &bookingOfflineUseCase{
		bData: bOCData,
	}
}

func (booku *bookingOfflineUseCase) GetListBookingOffline(data bookingOffline.OfflineClassUser) (list []bookingOffline.OfflineClassUser, err error) {
	offlineClass, err := booku.bData.SelectAllBookingOffline(data)
	if err != nil {
		return nil, err
	}
	return offlineClass, nil
}

func (booku *bookingOfflineUseCase) MemberBookingOffline(userID int, classID int) (err error) {
	if err := booku.bData.InsertMemberBookingOffline(userID, classID); err != nil {
		return err
	}
	return nil
}
