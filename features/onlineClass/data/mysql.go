package data

import (
	news "capstone/backend/features/News"
	"capstone/backend/features/onlineClass"
	"errors"
	"fmt"

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
	var record []News
	if err := nr.Conn.Preload("").Find(&record).Error; err != nil {
		return news
	}
	return toNewsCoreList(record)
}
func (or *mysqlOnlineClassRepo) SelectClassById(id int) (onlineClass.OnlineClassCore, error) {
	var idNews News

	err := nr.Conn.First(&idNews, id).Error

	if idNews.Title == "" && idNews.ID == 0 {
		return news.NewsCore{}, errors.New("News not found")
	}

	if err != nil {
		return news.NewsCore{}, err
	}

	return toNewsCore(idNews), nil
}
func (or *mysqlOnlineClassRepo) InsertClass(data onlineClass.OnlineClassCore) (err error) {
	record := fromCore(data)

	if err := nr.Conn.Create(&record).Error; err != nil {
		return news.NewsCore{}, err
	}

	return data, nil
}
func (or *mysqlOnlineClassRepo) UpdateClass(id int) (news onlineClass.OnlineClassCore, err error) {
	var singleNews News
	fmt.Println("Isi single account : ", singleNews)
	fmt.Println("id : ", id)
	err = nr.Conn.Model(&singleNews).Where("id=?", id).Updates(&singleNews).Error

	if err != nil {
		return news, err
	}

	return toNewsCore(singleNews), nil
}
func (or *mysqlOnlineClassRepo) DestryoClass(id int) (news onlineClass.OnlineClassCore, err error) {
	var singleNews News
	fmt.Println("Isi single account : ", singleNews)
	fmt.Println("id : ", id)
	err = nr.Conn.Model(&singleNews).Where("id=?", id).Updates(&singleNews).Error

	if err != nil {
		return news, err
	}

	return toNewsCore(singleNews), nil
}
