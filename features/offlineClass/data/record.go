package data

import (
	"capstone/backend/features/offlineClass"
	"time"
)

type OfflineClassCore struct {
	ID        int
	Name      string `json:"name"`
	Day       string `json:"day"`
	Date      string `json:"date"`
	Location  string `json:"location"`
	Time      string `json:"time"`
	Trainer   string `json:"trainer"`
	Image     string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toClassRecord(oc offlineClass.OfflineClassCore) OfflineClassCore {
	return OfflineClassCore{
		ID:        oc.ID,
		Name:      oc.Name,
		Day:       oc.Day,
		Date:      oc.Date,
		Location:  oc.Location,
		Time:      oc.Time,
		Trainer:   oc.Trainer,
		Image:     oc.Image,
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}
}

func toClassCore(oc OfflineClassCore) offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
		ID:        oc.ID,
		Name:      oc.Name,
		Day:       oc.Day,
		Date:      oc.Date,
		Location:  oc.Location,
		Time:      oc.Time,
		Trainer:   oc.Trainer,
		Image:     oc.Image,
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}
}

func toOfflineClassCoreList(oc []OfflineClassCore) []offlineClass.OfflineClassCore {
	convOC := []offlineClass.OfflineClassCore{}

	for _, ocList := range oc {
		convOC = append(convOC, toClassCore(ocList))
	}
	return convOC
}
