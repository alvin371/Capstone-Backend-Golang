package presentation

import (
	"capstone/backend/features/onlineClass"
	presentation_response "capstone/backend/features/onlineClass/presentation/rep"
	presentation_request "capstone/backend/features/onlineClass/presentation/req"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OnlineClassHandler struct {
	onlineClassBussiness onlineClass.Bussiness
}

func NewOnlineClassHandler(onlineClassBussiness onlineClass.Bussiness) *OnlineClassHandler {
	return &OnlineClassHandler{onlineClassBussiness}
}

func (oh *OnlineClassHandler) GetAllClassHandler(e echo.Context) error {
	data, err := oh.onlineClassBussiness.GetAllClass(onlineClass.OnlineClassCore{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.ToCoreSlice(data),
	})
}
func (oh *OnlineClassHandler) GetClassByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := oh.onlineClassBussiness.GetClassById(id)
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
func (oh *OnlineClassHandler) CreateClassHandler(e echo.Context) error {
	newOnlineClass := presentation_request.OnlineClass{}
	fmt.Println("testing", e)

	if err := e.Bind(&newOnlineClass); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := oh.onlineClassBussiness.CreateClass(newOnlineClass.ToClassCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newOnlineClass,
	})
}
func (oh *OnlineClassHandler) UpdateClassHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := oh.onlineClassBussiness.EditClass(id)
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
func (oh *OnlineClassHandler) DeleteClassHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := oh.onlineClassBussiness.DeleteClass(id)
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
