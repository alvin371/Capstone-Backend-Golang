package data

import (
	"capstone/backend/features/onlineClass"
	"time"
)

type OnlineClassCore struct {
	ID        int
	Name      string `json:"name"`
	Day       string `json:"day"`
	Date      string `json:"date"`
	Link      string `json:"link"`
	Time      string `json:"time"`
	Trainer   string `json:"trainer"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toClassRecord(oc onlineClass.OnlineClassCore) OnlineClassCore {
	return OnlineClassCore{
		ID:        oc.ID,
		Name:      oc.Name,
		Day:       oc.Day,
		Date:      oc.Date,
		Link:      oc.Link,
		Time:      oc.Time,
		Trainer:   oc.Trainer,
		Image:     oc.Image,
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}
}

func toClassCore(oc OnlineClassCore) onlineClass.OnlineClassCore {
	return onlineClass.OnlineClassCore{
		ID:        oc.ID,
		Name:      oc.Name,
		Day:       oc.Day,
		Date:      oc.Date,
		Link:      oc.Link,
		Time:      oc.Time,
		Trainer:   oc.Trainer,
		Image:     oc.Image,
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}
}

func toOnlineClassCoreList(oc []OnlineClassCore) []onlineClass.OnlineClassCore {
	convOC := []onlineClass.OnlineClassCore{}

	for _, ocList := range oc {
		convOC = append(convOC, toClassCore(ocList))
	}
	return convOC
}
