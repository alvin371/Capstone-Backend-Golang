package data

import (
	"capstone/backend/features/content/entity"
)

func (mysql *MySQL) Insert(data *entity.Video) error {
	// process insert data into db
	if err := mysql.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (mysql *MySQL) Update(id int, data *entity.Video) error {
	var video entity.Video

	if err := mysql.DB.Where("id = ?", id).First(&video).Error; err != nil {
		return err
	}

	// process insert data into db
	if err := mysql.DB.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (mysql *MySQL) Delete(id int) error {
	var video entity.Video

	if err := mysql.DB.Where("id = ?", id).Delete(&video).Error; err != nil {
		return err
	}

	return nil
}
