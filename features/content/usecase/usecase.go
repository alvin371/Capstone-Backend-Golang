package usecase

import (
	"capstone/backend/features/content"
	"capstone/backend/features/content/entity"

	"gorm.io/gorm"
)

type Service struct {
	DB          *gorm.DB
	contentRepo content.ContentRepository
}

func NewService(DB *gorm.DB, contentRepo content.ContentRepository) content.ContentUsecase {
	return &Service{
		DB:          DB,
		contentRepo: contentRepo,
	}
}

func (s *Service) CreateVideo(req entity.Video) error {
	err := s.contentRepo.Insert(&req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateVideo(req entity.Video, id int) error {
	err := s.contentRepo.Update(id, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteVideo(id int) error {
	err := s.contentRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() ([]*entity.Video, error) {
	videos, err := s.contentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *Service) GetVideosById(id int) (*entity.Video, error) {
	video, err := s.contentRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return video, nil
}
