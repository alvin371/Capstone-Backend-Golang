package data

import "capstone/backend/features/content/entity"

func (mysql *MySQL) GetAll() ([]*entity.Video, error) {
	var videos []*entity.Video

	if err := mysql.DB.Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func (mysql *MySQL) Get(id int) (*entity.Video, error) {
	var video *entity.Video

	if err := mysql.DB.First(&video, id).Error; err != nil {
		return nil, err
	}

	return video, nil
}
