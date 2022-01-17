package presentation

import (
	news "capstone/backend/features/News"
	"net/http"
	"strconv"

	presentation_response "capstone/backend/features/News/presentation/rep"
	presentation_request "capstone/backend/features/News/presentation/req"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsBussiness news.Bussiness
}

func NewNewsHandler(nd news.Bussiness) *NewsHandler {
	return &NewsHandler{
		newsBussiness: nd,
	}
}

func (nh *NewsHandler) GetNewsByIDHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Isi ID : ", id)
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
	result := nh.newsBussiness.GetAllNews("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All is Well",
		"data":    presentation_response.ToCoreSlice(result),
	})
}

func (nh *NewsHandler) CreateNewsHandler(e echo.Context) error {
	newNews := presentation_request.News{}
	e.Bind(newNews)
	resp, err := nh.newsBussiness.CreateNews(presentation_request.FromCore(newNews))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Data",
		"data":    presentation_response.ToCore(resp),
	})
}
func (nh *NewsHandler) EditNewsHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Test : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
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
