package content

import (
	"capstone/backend/features/content/entity"
)

type ContentRepository interface {
	Insert(data *entity.Video) error
	Update(id int, data *entity.Video) error
	Delete(id int) error
	GetAll() ([]*entity.Video, error)
	Get(id int) (*entity.Video, error)
}
