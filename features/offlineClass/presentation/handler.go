package presentation

import (
	"capstone/backend/features/offlineClass"
	presentation_response "capstone/backend/features/offlineClass/presentation/rep"
	presentation_request "capstone/backend/features/offlineClass/presentation/req"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OfflineClassHandler struct {
	offlineClassBussiness offlineClass.Bussiness
}

func NewOfflineClassHandler(offlineClassBussiness offlineClass.Bussiness) *OfflineClassHandler {
	return &OfflineClassHandler{offlineClassBussiness}
}

func (oh *OfflineClassHandler) GetAllOfflineClassHandler(e echo.Context) error {
	data, err := oh.offlineClassBussiness.GetAllClass(offlineClass.OfflineClassCore{})

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
func (oh *OfflineClassHandler) GetOfflineClassByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Isi ID : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := oh.offlineClassBussiness.GetClassById(id)
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
func (oh *OfflineClassHandler) CreateOfflineClassHandler(e echo.Context) error {
	newOnlineClass := presentation_request.OfflineClassCore{}

	if err := e.Bind(&newOnlineClass); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("testing", newOnlineClass)

	if err := oh.offlineClassBussiness.CreateClass(newOnlineClass.ToClassCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newOnlineClass,
	})
}
func (oh *OfflineClassHandler) UpdateOfflineClassHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	updateOnlineClass := presentation_request.OfflineClassCore{}
	if err := e.Bind(&updateOnlineClass); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("Test : ", id)
	fmt.Println("testing", updateOnlineClass)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := oh.offlineClassBussiness.EditClass(id, updateOnlineClass.ToClassCore())
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
func (oh *OfflineClassHandler) DeleteOfflineClassHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Test : ", id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "id tidak ditemukan",
		})
	}
	data, err := oh.offlineClassBussiness.DeleteClass(id)
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
