package factory

import (
	"capstone/backend/driver"
	_news_bussiness "capstone/backend/features/News/bussiness"
	_news_data "capstone/backend/features/News/data"
	_news_presentation "capstone/backend/features/News/presentation"

	_onlineClass_bussiness "capstone/backend/features/onlineClass/bussiness"
	_onlineClass_data "capstone/backend/features/onlineClass/data"
	_onlineClass_presentation "capstone/backend/features/onlineClass/presentation"
)

type Presenter struct {
	NewsPresentation        *_news_presentation.NewsHandler
	OnlineClassPresentation *_onlineClass_presentation.OnlineClassHandler
}

func Init() Presenter {

	newsData := _news_data.NewNewsRepository(driver.DB)
	newsBussiness := _news_bussiness.NewBussinessNews(newsData)
	newsPresentation := _news_presentation.NewNewsHandler(newsBussiness)

	onlineClassData := _onlineClass_data.NewOnlineClassRepository(driver.DB)
	onlineClassBussiness := _onlineClass_bussiness.NewBussinessOnlineClass(onlineClassData)
	onlineClassPresentation := _onlineClass_presentation.NewOnlineClassHandler(onlineClassBussiness)
	return Presenter{
		NewsPresentation:        newsPresentation,
		OnlineClassPresentation: onlineClassPresentation,
	}

}
