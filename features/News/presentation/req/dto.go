package req

import (
	news "capstone/backend/features/News"
)

type News struct {
	Title       string `json: "title" header:"title"`
	Description string `json: "description" header:"description"`
	Content     string `json: "content" header:"content"`
	CreatorName string `json: "creator" header:"creator"`
	Picture     string `json: "picture" header:"picture"`
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
