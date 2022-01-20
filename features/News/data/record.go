package data

import (
	news "capstone/backend/features/News"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title       string `json: "title"`
	Description string `json: "description"`
	Content     string `json: "content"`
	CreatorName string `json: "creator"`
	Picture     string `json: "picture"`
}

func toNewsRecord(news news.NewsCore) News {
	return News{
		Model: gorm.Model{
			ID: uint(news.ID),
		},
		Title:       news.Title,
		Description: news.Description,
		Content:     news.Content,
		CreatorName: news.CreatorName,
		Picture:     news.Picture,
	}
}

func toNewsCore(nws News) news.NewsCore {
	return news.NewsCore{
		ID:          int(nws.ID),
		Title:       nws.Title,
		Description: nws.Description,
		Content:     nws.Content,
		CreatorName: nws.CreatorName,
		Picture:     nws.Picture,
	}
}

func toNewsCoreList(nws []News) []news.NewsCore {
	convNews := []news.NewsCore{}

	for _, newslist := range nws {
		convNews = append(convNews, toNewsCore(newslist))
	}
	return convNews
}

func fromCore(core news.NewsCore) News {
	return News{}
}
