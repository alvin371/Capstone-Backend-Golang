package entity

import "time"

type Video struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Duration  string     `json:"duration"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
