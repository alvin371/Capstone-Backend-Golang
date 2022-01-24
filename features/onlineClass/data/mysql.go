package data

import (
	"capstone/backend/features/onlineClass"
	"errors"

	"gorm.io/gorm"
)

type mysqlOnlineClassRepo struct {
	Conn *gorm.DB
}

func NewOnlineClassRepository(conn *gorm.DB) onlineClass.Data {
	return &mysqlOnlineClassRepo{
		Conn: conn,
	}
}

func (or *mysqlOnlineClassRepo) SelectAllClass(onlineClass.OnlineClassCore) (class []onlineClass.OnlineClassCore, err error) {
	var record []OnlineClass
	err = or.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toOnlineClassCoreList(record), nil
}
func (or *mysqlOnlineClassRepo) SelectClassById(id int) (onlineClass.OnlineClassCore, error) {
	var idOclass OnlineClass

	err := or.Conn.First(&idOclass, id).Error

	if idOclass.Name == "" && idOclass.ID == 0 {
		return onlineClass.OnlineClassCore{}, errors.New("News not found")
	}

	if err != nil {
		return onlineClass.OnlineClassCore{}, err
	}

	return toClassCore(idOclass), nil
}
func (or *mysqlOnlineClassRepo) InsertClass(data onlineClass.OnlineClassCore) (err error) {
	convData := toClassRecord(data)

	if err := or.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
func (or *mysqlOnlineClassRepo) UpdateClass(id int) (news onlineClass.OnlineClassCore, err error) {
	var singleClass OnlineClass
	// fmt.Println("Isi single account : ", singleClass)
	// fmt.Println("id : ", id)
	err = or.Conn.Model(&singleClass).Where("id=?", id).Updates(&singleClass).Error

	if err != nil {
		return news, err
	}

	return toClassCore(singleClass), nil
}
func (or *mysqlOnlineClassRepo) DestryoClass(id int) (news onlineClass.OnlineClassCore, err error) {
	var singleClass OnlineClass
	// fmt.Println("Isi single account : ", singleClass)
	// fmt.Println("id : ", id)
	err = or.Conn.Model(&singleClass).Where("id=?", id).Delete(&singleClass).Error

	if err != nil {
		return news, err
	}

	return toClassCore(singleClass), nil
}
