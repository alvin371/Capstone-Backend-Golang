package rep

import "capstone/backend/features/onlineClass"

type OnlineClassCore struct {
	Name    string `json:"name"`
	Day     string `json:"day"`
	Date    string `json:"date"`
	Link    string `json:"link"`
	Time    string `json:"time"`
	Trainer string `json:"trainer"`
	Image   string `json:"image"`
}

func ToCore(req onlineClass.OnlineClassCore) OnlineClassCore {
	return OnlineClassCore{
		Name:    req.Name,
		Day:     req.Day,
		Date:    req.Date,
		Link:    req.Link,
		Time:    req.Time,
		Trainer: req.Trainer,
		Image:   req.Image,
	}
}

func ToCoreSlice(core []onlineClass.OnlineClassCore) []OnlineClassCore {
	var onlineClassArray []OnlineClassCore
	for key := range core {
		onlineClassArray = append(onlineClassArray, ToCore(core[key]))
	}
	return onlineClassArray
}
