package data

import "capstone/backend/features/offline-class/entity"

func (mysql *MySQL) GetAll() ([]*entity.OfflineClass, error) {
	var class []*entity.OfflineClass

	if err := mysql.DB.Find(&class).Error; err != nil {
		return nil, err
	}

	return class, nil
}

func (mysql *MySQL) Get(id int) (*entity.OfflineClass, error) {
	var class *entity.OfflineClass

	if err := mysql.DB.First(&class, id).Error; err != nil {
		return nil, err
	}

	return class, nil
}
