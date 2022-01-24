package news

import "time"

type NewsCore struct {
	ID          int       `param:"id"`
	Title       string    `json: "title" param:"title"`
	Description string    `json: "description" param:"description"`
	Content     string    `json: "content" param:"content"`
	CreatorName string    `json: "creator_name" param:"creator_name"`
	Picture     string    `json: "picture" param:"picture"`
	Created_at  time.Time `json: "created_at" `
	Updated_at  time.Time `json: "updated_at" `
}

type Bussiness interface {
	GetNewsByID(id int) (NewsCore, error)
	GetAllNews(NewsCore) (news []NewsCore, err error)
	CreateNews(data NewsCore) (err error)
	EditNews(id int) (news NewsCore, err error)
}

type Data interface {
	SelectIdNews(id int) (NewsCore, error)
	SelectNewsAll(NewsCore) (news []NewsCore, err error)
	InsertNews(data NewsCore) (err error)
	UpdateNews(id int) (news NewsCore, err error)
}
