package content

import "capstone/backend/features/content/entity"

type ContentUsecase interface {
	CreateVideo(req entity.Video) error
	UpdateVideo(req entity.Video, id int) error
	DeleteVideo(id int) error
	GetAll() ([]*entity.Video, error)
	GetVideosById(id int) (*entity.Video, error)
}
