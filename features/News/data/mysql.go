package data

import (
	news "capstone/backend/features/News"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type mysqlNewsRepository struct {
	Conn *gorm.DB
}

func NewNewsRepository(conn *gorm.DB) news.Data {
	return &mysqlNewsRepository{
		Conn: conn,
	}
}

func (nr *mysqlNewsRepository) SelectIdNews(id int) (news.NewsCore, error) {
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
func (nr *mysqlNewsRepository) SelectNewsAll(search string) (news []news.NewsCore) {
	var record []News
	if err := nr.Conn.Preload("").Find(&record).Error; err != nil {
		return news
	}
	return toNewsCoreList(record)
}
func (nr *mysqlNewsRepository) InsertNews(data news.NewsCore) (resp news.NewsCore, err error) {
	record := fromCore(data)

	if err := nr.Conn.Create(&record).Error; err != nil {
		return news.NewsCore{}, err
	}

	return data, nil
}
func (nr *mysqlNewsRepository) UpdateNews(id int) (news news.NewsCore, err error) {
	var singleNews News
	fmt.Println("Isi single account : ", singleNews)
	fmt.Println("id : ", id)
	err = nr.Conn.Model(&singleNews).Where("id=?", id).Updates(&singleNews).Error

	if err != nil {
		return news, err
	}

	return toNewsCore(singleNews), nil
}
