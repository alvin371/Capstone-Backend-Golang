package data

import (
	"capstone/backend/features/offlineClass"
	"errors"

	"gorm.io/gorm"
)

type mysqlOfflineClassRepo struct {
	Conn *gorm.DB
}

func NewofflineClassRepository(conn *gorm.DB) offlineClass.Data {
	return &mysqlOfflineClassRepo{
		Conn: conn,
	}
}

func (or *mysqlOfflineClassRepo) SelectAllClass(offlineClass.OfflineClassCore) (class []offlineClass.OfflineClassCore, err error) {
	var record []OfflineClassCore
	err = or.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toOfflineClassCoreList(record), nil
}
func (or *mysqlOfflineClassRepo) SelectClassById(id int) (offlineClass.OfflineClassCore, error) {
	var idOclass OfflineClassCore

	err := or.Conn.First(&idOclass, id).Error

	if idOclass.Name == "" && idOclass.ID == 0 {
		return offlineClass.OfflineClassCore{}, errors.New("News not found")
	}

	if err != nil {
		return offlineClass.OfflineClassCore{}, err
	}

	return toClassCore(idOclass), nil
}
func (or *mysqlOfflineClassRepo) InsertClass(data offlineClass.OfflineClassCore) (err error) {
	convData := toClassRecord(data)

	if err := or.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
func (or *mysqlOfflineClassRepo) UpdateClass(id int, data offlineClass.OfflineClassCore) (news offlineClass.OfflineClassCore, err error) {
	var singleClass OfflineClassCore
	// fmt.Println("Isi single account : ", singleClass)
	// fmt.Println("id : ", id)
	err = or.Conn.Model(&singleClass).Where("id=?", id).Updates(&data).Error

	if err != nil {
		return news, err
	}

	return toClassCore(singleClass), nil
}
func (or *mysqlOfflineClassRepo) DestroyClass(id int) (news offlineClass.OfflineClassCore, err error) {
	var singleClass OfflineClassCore
	// fmt.Println("Isi single account : ", singleClass)
	// fmt.Println("id : ", id)
	err = or.Conn.Model(&singleClass).Where("id=?", id).Delete(&singleClass).Error

	if err != nil {
		return news, err
	}

	return toClassCore(singleClass), nil
}
