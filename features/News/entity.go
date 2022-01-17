package news

import "time"

type NewsCore struct {
	ID          int
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Content     string    `json: "content"`
	CreatorName string    `json: "creator"`
	Picture     string    `json: "picture"`
	Created_at  time.Time `json: "created_at"`
	Updated_at  time.Time `json: "updated_at"`
}

type Bussiness interface {
	GetNewsByID(id int) (NewsCore, error)
	GetAllNews(search string) (news []NewsCore)
	CreateNews(data NewsCore) (resp NewsCore, err error)
	EditNews(id int) (news NewsCore, err error)
}

type Data interface {
	SelectIdNews(id int) (NewsCore, error)
	SelectNewsAll(search string) (news []NewsCore)
	InsertNews(data NewsCore) (resp NewsCore, err error)
	UpdateNews(id int) (news NewsCore, err error)
}
