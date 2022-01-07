package usecase

import (
	offline_class "capstone/backend/features/offline-class"
	"capstone/backend/features/offline-class/entity"

	"gorm.io/gorm"
)

type Service struct {
	DB        *gorm.DB
	classRepo offline_class.OfflineClassRepository
}

func NewService(DB *gorm.DB, classRepo offline_class.OfflineClassRepository) offline_class.OfflineClassUsecase {
	return &Service{
		DB:        DB,
		classRepo: classRepo,
	}
}

func (s *Service) CreateClass(req entity.OfflineClass) error {
	err := s.classRepo.Insert(&req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateClass(req entity.OfflineClass, id int) error {
	err := s.classRepo.Update(id, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteClass(id int) error {
	err := s.classRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() ([]*entity.OfflineClass, error) {
	classes, err := s.classRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (s *Service) GetClassById(id int) (*entity.OfflineClass, error) {
	class, err := s.classRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return class, nil
}
