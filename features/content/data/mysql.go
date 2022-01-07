package data

import (
	"capstone/backend/features/content"

	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(db *gorm.DB) content.ContentRepository {
	return &MySQL{
		DB: db,
	}
}
