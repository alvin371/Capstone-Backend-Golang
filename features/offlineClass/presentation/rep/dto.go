package rep

import "capstone/backend/features/offlineClass"

type OfflineClassCore struct {
	Name     string `json:"name"`
	Day      string `json:"day"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Time     string `json:"time"`
	Trainer  string `json:"trainer"`
	Image    string `json:"image"`
}

func ToCore(req offlineClass.OfflineClassCore) OfflineClassCore {
	return OfflineClassCore{
		Name:     req.Name,
		Day:      req.Day,
		Date:     req.Date,
		Location: req.Location,
		Time:     req.Time,
		Trainer:  req.Trainer,
		Image:    req.Image,
	}
}

func ToCoreSlice(core []offlineClass.OfflineClassCore) []OfflineClassCore {
	var onlineClassArray []OfflineClassCore
	for key := range core {
		onlineClassArray = append(onlineClassArray, ToCore(core[key]))
	}
	return onlineClassArray
}
