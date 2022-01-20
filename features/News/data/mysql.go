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
func (nr *mysqlNewsRepository) SelectNewsAll(news.NewsCore) ([]news.NewsCore, error) {
	var record []News
	err := nr.Conn.Find(&record).Error

	if err != nil {
		return nil, err
	}
	return toNewsCoreList(record), nil
}
func (nr *mysqlNewsRepository) InsertNews(data news.NewsCore) (err error) {
	convData := toNewsRecord(data)

	if err := nr.Conn.Create(&convData).Error; err != nil {
		return err
	}
	return nil
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
