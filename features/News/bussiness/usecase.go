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
func (nu *NewsUsecase) GetAllNews(data news.NewsCore) (resp []news.NewsCore, err error) {
	news, err := nu.nData.SelectNewsAll(data)
	if err != nil {
		return nil, err
	}
	return news, nil
}
func (nu *NewsUsecase) CreateNews(data news.NewsCore) (err error) {
	if err := nu.nData.InsertNews(data); err != nil {
		return err
	}
	return nil
}
func (nu *NewsUsecase) EditNews(id int) (news news.NewsCore, err error) {
	newsData, err := nu.nData.UpdateNews(id)

	if err != nil {
		return news, err
	}
	return newsData, nil
}
