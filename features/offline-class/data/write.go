package data

import (
	"capstone/backend/features/offline-class/entity"
)

func (mysql *MySQL) Insert(data *entity.OfflineClass) error {
	// process insert data into db
	if err := mysql.DB.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (mysql *MySQL) Update(id int, data *entity.OfflineClass) error {
	var class entity.OfflineClass

	if err := mysql.DB.Where("id = ?", id).First(&class).Error; err != nil {
		return err
	}

	// process insert data into db
	if err := mysql.DB.Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (mysql *MySQL) Delete(id int) error {
	var class entity.OfflineClass

	if err := mysql.DB.Where("id = ?", id).Delete(&class).Error; err != nil {
		return err
	}

	return nil
}
