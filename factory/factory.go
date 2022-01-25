package factory

import (
	"capstone/backend/driver"
	_news_bussiness "capstone/backend/features/News/bussiness"
	_news_data "capstone/backend/features/News/data"
	_news_presentation "capstone/backend/features/News/presentation"

	_user_bussiness "capstone/backend/features/User/bussiness"
	_user_data "capstone/backend/features/User/data"
	_user_presentation "capstone/backend/features/User/presentation"

	_onlineClass_bussiness "capstone/backend/features/onlineClass/bussiness"
	_onlineClass_data "capstone/backend/features/onlineClass/data"
	_onlineClass_presentation "capstone/backend/features/onlineClass/presentation"

	_offlineClass_bussiness "capstone/backend/features/offlineClass/bussiness"
	_offlineClass_data "capstone/backend/features/offlineClass/data"
	_offlineClass_presentation "capstone/backend/features/offlineClass/presentation"

	_bookingOffline_bussiness "capstone/backend/features/bookingOffline/bussiness"
	_bookingOffline_data "capstone/backend/features/bookingOffline/data"
	_bookingOffline_presentation "capstone/backend/features/bookingOffline/presentation"
)

type Presenter struct {
	NewsPresentation                  *_news_presentation.NewsHandler
	UserPresentation                  *_user_presentation.UserHandler
	OnlineClassPresentation           *_onlineClass_presentation.OnlineClassHandler
	PresenterOfflineClassPresentation *_offlineClass_presentation.OfflineClassHandler
	BookingOfflinePresentation        *_bookingOffline_presentation.BookingOfflineHandler
}

func Init() Presenter {

	newsData := _news_data.NewNewsRepository(driver.DB)
	newsBussiness := _news_bussiness.NewBussinessNews(newsData)
	newsPresentation := _news_presentation.NewNewsHandler(newsBussiness)

	// user
	userData := _user_data.NewMySqlUSer(driver.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewHandlerAccount(userBussiness)

	onlineClassData := _onlineClass_data.NewOnlineClassRepository(driver.DB)
	onlineClassBussiness := _onlineClass_bussiness.NewBussinessOnlineClass(onlineClassData)
	onlineClassPresentation := _onlineClass_presentation.NewOnlineClassHandler(onlineClassBussiness)

	offlineClassData := _offlineClass_data.NewofflineClassRepository(driver.DB)
	offlineClassBussiness := _offlineClass_bussiness.NewBussinessOfflineClass(offlineClassData)
	offlineClassPresentation := _offlineClass_presentation.NewOfflineClassHandler(offlineClassBussiness)

	bookingOfflineData := _bookingOffline_data.NewBookingOfflineRepository(driver.DB)
	bookingOfflineBussiness := _bookingOffline_bussiness.NewBussinessBookingOfflineClass(bookingOfflineData)
	bookingOfflinePresentation := _bookingOffline_presentation.NewBookingOfflineHandler(bookingOfflineBussiness)

	return Presenter{
		NewsPresentation:                  newsPresentation,
		UserPresentation:                  userPresentation,
		OnlineClassPresentation:           onlineClassPresentation,
		PresenterOfflineClassPresentation: offlineClassPresentation,
		BookingOfflinePresentation:        &bookingOfflinePresentation,
	}

}
