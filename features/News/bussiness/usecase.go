package bussiness

import (
	news "capstone/backend/features/News"

	"github.com/go-playground/validator/v10"
)

type NewsUsecase struct {
	nData    news.Data
	validate *validator.Validate
}

func NewBussinessNews(newsData news.Data) news.Bussiness {
	return &NewsUsecase{
		nData:    newsData,
		validate: validator.New(),
	}
}

func (nu *NewsUsecase) GetNewsByID(id int) (news.NewsCore, error) {
	newsData, err := nu.nData.SelectIdNews(id)

	if err != nil {
		return news.NewsCore{}, err
	}
	return newsData, nil
}
func (nu *NewsUsecase) GetAllNews(search string) (resp []news.NewsCore) {
	resp = nu.nData.SelectNewsAll(search)
	return
}
func (nu *NewsUsecase) CreateNews(data news.NewsCore) (resp news.NewsCore, err error) {
	if err := nu.validate.Struct(data); err != nil {
		return news.NewsCore{}, err
	}

	resp, err = nu.nData.InsertNews(data)
	if err != nil {
		return news.NewsCore{}, err
	}
	return news.NewsCore{}, nil
}
func (nu *NewsUsecase) EditNews(id int) (news news.NewsCore, err error) {
	newsData, err := nu.nData.UpdateNews(id)

	if err != nil {
		return news, err
	}
	return newsData, nil
}
