package req

import "capstone/backend/features/onlineClass"

type OnlineClassCore struct {
	Name    string `json:"name" form:"name"`
	Day     string `json:"day" form:"day"`
	Date    string `json:"date" form:"date"`
	Link    string `json:"link" form:"link"`
	Time    string `json:"time" form:"time"`
	Trainer string `json:"trainer" form:"trainer"`
	Image   string `json:"image" form:"image"`
}

func FromCore(core OnlineClassCore) onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		Name:    core.Name,
		Day:     core.Day,
		Date:    core.Date,
		Link:    core.Link,
		Time:    core.Time,
		Trainer: core.Trainer,
		Image:   core.Image,
	}
}

func (core *OnlineClassCore) ToClassCore() onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		Name:    core.Name,
		Day:     core.Day,
		Date:    core.Date,
		Link:    core.Link,
		Time:    core.Time,
		Trainer: core.Trainer,
		Image:   core.Image,
	}
}
