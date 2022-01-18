package factory

import (
	"capstone/backend/driver"
	_news_bussiness "capstone/backend/features/News/bussiness"
	_news_data "capstone/backend/features/News/data"
	_news_presentation "capstone/backend/features/News/presentation"

	_user_bussiness "capstone/backend/features/User/bussiness"
	_user_data "capstone/backend/features/User/data"
	_user_presentation "capstone/backend/features/User/presentation"
)

type Presenter struct {
	NewsPresentation *_news_presentation.NewsHandler
	UserPresentation *_user_presentation.UserHandler
}

func Init() Presenter {

	newsData := _news_data.NewNewsRepository(driver.DB)
	newsBussiness := _news_bussiness.NewBussinessNews(newsData)
	newsPresentation := _news_presentation.NewNewsHandler(newsBussiness)

	// user
	userData := _user_data.NewMySqlUSer(driver.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewHandlerAccount(userBussiness)
	return Presenter{
		NewsPresentation: newsPresentation,
		UserPresentation: userPresentation,
	}

}
