package req

import "capstone/backend/features/onlineClass"

type OnlineClass struct {
	Name    string `json:"name"`
	Day     string `json:"day"`
	Date    string `json:"date"`
	Link    string `json:"link"`
	Time    string `json:"time"`
	Trainer string `json:"trainer"`
	Image   string `json:"image"`
}

func FromCore(core OnlineClass) onlineClass.OnlineClassCore {
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

func (core *OnlineClass) ToClassCore() onlineClass.OnlineClassCore {
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
