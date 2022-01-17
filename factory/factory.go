package factory

import (
	"capstone/backend/driver"
	_news_bussiness "capstone/backend/features/News/bussiness"
	_news_data "capstone/backend/features/News/data"
	_news_presentation "capstone/backend/features/News/presentation"
)

type Presenter struct {
	NewsPresentation *_news_presentation.NewsHandler
}

func Init() Presenter {

	newsData := _news_data.NewNewsRepository(driver.DB)
	newsBussiness := _news_bussiness.NewBussinessNews(newsData)
	newsPresentation := _news_presentation.NewNewsHandler(newsBussiness)
	return Presenter{
		NewsPresentation: newsPresentation,
	}

}
