package bussiness

import (
	news "capstone/backend/features/News"
	mocks "capstone/backend/features/News/mock"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData      mocks.Data
	newsBussiness news.Bussiness
	newsDatas     []news.NewsCore
	newsData      news.NewsCore
	newsLogin     news.NewsCore
)

func TestMain(m *testing.M) {
	newsBussiness = NewBussinessNews(&mockData)

	newsDatas = []news.NewsCore{
		{
			ID:          1,
			Title:       "pertama",
			Description: "Apa aja dehh",
			Content:     "Kok Bisa siiih",
			CreatorName: "alvin",
			Picture:     "local",
		},
	}

	newsData = news.NewsCore{
		Title:       "pertama",
		Description: "Apa aja dehh",
		Content:     "Kok Bisa siiih",
		CreatorName: "alvin",
		Picture:     "local",
	}
	os.Exit(m.Run())
}

func TestCreateNews(t *testing.T) {
	t.Run("validate create account", func(t *testing.T) {
		mockData.On("InsertAccount", mock.AnythingOfType("news.NewsCore")).Return(nil).Once()
		err := newsBussiness.CreateNews(news.NewsCore{})
		assert.Nil(t, err)
	})
	t.Run("error create account", func(t *testing.T) {
		mockData.On("InsertAccount", mock.AnythingOfType("news.NewsCore")).Return(errors.New("Error")).Once()
		err := newsBussiness.CreateNews(news.NewsCore{})
		assert.NotNil(t, err)
	})
}
func TestGetAllNews(t *testing.T) {
	t.Run("validate get accounts", func(t *testing.T) {
		mockData.On("SelectAccount", mock.AnythingOfType("news.NewsCore")).Return(newsDatas, nil).Once()
		resp, err := newsBussiness.GetAllNews(news.NewsCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get accounts", func(t *testing.T) {
		mockData.On("SelectAccount", mock.AnythingOfType("news.NewsCore")).Return(nil, errors.New("error"))
		resp, err := newsBussiness.GetAllNews(news.NewsCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestGetNewsByID(t *testing.T) {
	t.Run("validate get account by id", func(t *testing.T) {
		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(news.NewsCore{}, nil).Once()
		resp, err := newsBussiness.GetNewsByID(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("error get account by id", func(t *testing.T) {
		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(news.NewsCore{}, errors.New("error")).Once()
		resp, err := newsBussiness.GetNewsByID(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}
