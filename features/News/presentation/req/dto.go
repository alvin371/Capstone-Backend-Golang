package req

import (
	news "capstone/backend/features/News"
	"time"
)

type News struct {
	ID          int
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Content     string    `json: "content"`
	CreatorName string    `json: "creator"`
	Picture     string    `json: "picture"`
	Created_at  time.Time `json: "created_at"`
	Updated_at  time.Time `json: "updated_at"`
}

func FromCore(core News) news.NewsCore {
	return news.NewsCore{
		ID:          core.ID,
		Title:       core.Title,
		Description: core.Description,
		Content:     core.Content,
		CreatorName: core.CreatorName,
		Picture:     core.Picture,
		Created_at:  core.Created_at,
		Updated_at:  core.Updated_at,
	}
}
