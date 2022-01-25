package req

import "capstone/backend/features/offlineClass"

type OfflineClassCore struct {
	Name     string `json:"name" form:"name"`
	Day      string `json:"day" form:"day"`
	Date     string `json:"date" form:"date"`
	Location string `json:"location" form:"location"`
	Time     string `json:"time" form:"time"`
	Trainer  string `json:"trainer" form:"trainer"`
	Image    string `json:"image" form:"image"`
}

func FromCore(core OfflineClassCore) offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
		Name:     core.Name,
		Day:      core.Day,
		Date:     core.Date,
		Location: core.Location,
		Time:     core.Time,
		Trainer:  core.Trainer,
		Image:    core.Image,
	}
}

func (core *OfflineClassCore) ToClassCore() offlineClass.OfflineClassCore {
	return offlineClass.OfflineClassCore{
		Name:     core.Name,
		Day:      core.Day,
		Date:     core.Date,
		Location: core.Location,
		Time:     core.Time,
		Trainer:  core.Trainer,
		Image:    core.Image,
	}
}
