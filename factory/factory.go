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
)

type Presenter struct {
	NewsPresentation        *_news_presentation.NewsHandler
	OnlineClassPresentation *_onlineClass_presentation.OnlineClassHandler
	UserPresentation        *_user_presentation.UserHandler
}

func Init() Presenter {

	newsData := _news_data.NewNewsRepository(driver.DB)
	newsBussiness := _news_bussiness.NewBussinessNews(newsData)
	newsPresentation := _news_presentation.NewNewsHandler(newsBussiness)

	onlineClassData := _onlineClass_data.NewOnlineClassRepository(driver.DB)
	onlineClassBussiness := _onlineClass_bussiness.NewBussinessOnlineClass(onlineClassData)
	onlineClassPresentation := _onlineClass_presentation.NewOnlineClassHandler(onlineClassBussiness)

	// user
	userData := _user_data.NewMySqlUSer(driver.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewHandlerAccount(userBussiness)
	return Presenter{
		OnlineClassPresentation: onlineClassPresentation,
		NewsPresentation:        newsPresentation,
		UserPresentation:        userPresentation,
	}

}
