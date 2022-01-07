package offline_class

import (
	"capstone/backend/features/offline-class/entity"
)

type OfflineClassRepository interface {
	Insert(data *entity.OfflineClass) error
	Update(id int, data *entity.OfflineClass) error
	Delete(id int) error
	GetAll() ([]*entity.OfflineClass, error)
	Get(id int) (*entity.OfflineClass, error)
}
