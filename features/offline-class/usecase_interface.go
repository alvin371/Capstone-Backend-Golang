package offline_class

import "capstone/backend/features/offline-class/entity"

type OfflineClassUsecase interface {
	CreateClass(req entity.OfflineClass) error
	UpdateClass(req entity.OfflineClass, id int) error
	DeleteClass(id int) error
	GetAll() ([]*entity.OfflineClass, error)
	GetClassById(id int) (*entity.OfflineClass, error)
}
