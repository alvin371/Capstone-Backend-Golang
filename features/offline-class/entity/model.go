package entity

import "time"

type OfflineClass struct {
	ID               int        `json:"id"`
	TitleClass       string     `json:"title_class"`
	TrainerName      string     `json:"trainer_name"`
	DescriptionClass string     `json:"description_class"`
	Access           string     `json:"access"`
	Duration         string     `json:"duration"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}
