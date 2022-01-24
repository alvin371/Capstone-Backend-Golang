package onlineClass

import "time"

type OnlineClassCore struct {
	ID        int
	Name      string `json:"name"`
	Day       string `json:"day"`
	Date      string `json:"date`
	Link      string `json:"link"`
	Time      string `json:"time"`
	Trainer   string `json:"trainer`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	// crud
	GetAllClass(OnlineClassCore) (class []OnlineClassCore)
	GetClassById(id int) (OnlineClassCore, error)
	CreateClass(data OnlineClassCore) (resp OnlineClassCore, err error)
	EditClass(id int) (news OnlineClassCore, err error)
	DeleteClass(id int) (news OnlineClassCore, err error)
}

type Data interface {
	SelectAllClass(OnlineClassCore) (class []OnlineClassCore)
	SelectClassById(id int) (OnlineClassCore, error)
	InsertClass(data OnlineClassCore) (resp OnlineClassCore, err error)
	UpdateClass(id int) (news OnlineClassCore, err error)
	DestryoClass(id int) (news OnlineClassCore, err error)
}
