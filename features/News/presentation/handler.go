package presentation

import (
	news "capstone/backend/features/News"
	"fmt"
	"net/http"
	"strconv"

	presentation_response "capstone/backend/features/News/presentation/rep"
	presentation_request "capstone/backend/features/News/presentation/req"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsBussiness news.Bussiness
}

func NewNewsHandler(newsBussiness news.Bussiness) *NewsHandler {
	return &NewsHandler{newsBussiness}
}

func (nh *NewsHandler) GetNewsByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := nh.newsBussiness.GetNewsByID(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCore(data),
	})
}
func (nh *NewsHandler) GetAllNewsHandler(c echo.Context) error {
	data, err := nh.newsBussiness.GetAllNews(news.NewsCore{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCoreSlice(data),
	})

}

func (nh *NewsHandler) CreateNewsHandler(c echo.Context) error {
	newNews := presentation_request.News{}
	fmt.Println("testing", c)

	if err := c.Bind(&newNews); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := nh.newsBussiness.CreateNews(newNews.ToNewsCore()); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newNews,
	})
}
func (nh *NewsHandler) EditNewsHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := nh.newsBussiness.EditNews(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCore(data),
	})
}
