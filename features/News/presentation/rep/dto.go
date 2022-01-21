package rep

import news "capstone/backend/features/News"

type News struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	Content     string `json: "content"`
	CreatorName string `json: "creator_name"`
	Picture     string `json: "picture"`
}

func ToCore(req news.NewsCore) News {
	return News{
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		CreatorName: req.CreatorName,
		Picture:     req.Picture,
	}
}

func ToCoreSlice(core []news.NewsCore) []News {
	var newsArray []News
	for key := range core {
		newsArray = append(newsArray, ToCore(core[key]))
	}
	return newsArray
}
