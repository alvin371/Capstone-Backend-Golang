package data

import (
	offline_class "capstone/backend/features/offline-class"

	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(db *gorm.DB) offline_class.OfflineClassRepository {
	return &MySQL{
		DB: db,
	}
}
