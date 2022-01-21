package req

import (
	news "capstone/backend/features/News"
)

type News struct {
	Title       string `json: "title" param:"title"`
	Description string `json: "description" param:"description"`
	Content     string `json: "content" param:"content"`
	CreatorName string `json: "creator_name" param:"creator_name"`
	Picture     string `json: "picture" param:"picture"`
}

func FromCore(core News) news.NewsCore {
	return news.NewsCore{
		Title:       core.Title,
		Description: core.Description,
		Content:     core.Content,
		CreatorName: core.CreatorName,
		Picture:     core.Picture,
	}
}
func (core *News) ToNewsCore() news.NewsCore {
	return news.NewsCore{
		Title:       core.Title,
		Content:     core.Content,
		Description: core.Description,
		CreatorName: core.CreatorName,
		Picture:     core.Picture,
	}
}
