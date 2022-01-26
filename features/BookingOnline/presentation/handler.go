package presentation

import (
	"capstone/backend/features/bookingOnline"
	presentation_response "capstone/backend/features/bookingOnline/presentation/rep"
	presentation_request "capstone/backend/features/bookingOnline/presentation/req"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingOnlineHandler struct {
	bookingOfflineBussiness bookingOnline.Bussiness
}

func NewBookingOnlineHandler(bookingOfflineBussiness bookingOnline.Bussiness) BookingOnlineHandler {
	return BookingOnlineHandler{bookingOfflineBussiness}
}

func (boh *BookingOnlineHandler) GetListBookingOnline(e echo.Context) error {
	data, err := boh.bookingOfflineBussiness.GetListBookingOnline(bookingOnline.OnlineClassUser{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get List Booking Offline",
		"data":    presentation_response.TobookingOnlineCoreList(data),
	})
}

func (boh *BookingOnlineHandler) InsertMemberBookingOnline(e echo.Context) error {
	newBookingOnline := presentation_request.OnlineClassUser{}

	if err := e.Bind(&newBookingOnline); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	fmt.Println("testing", newBookingOnline)

	if err := boh.bookingOfflineBussiness.MemberBookingOnline(newBookingOnline.ClassID, newBookingOnline.UserID); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newBookingOnline,
	})
}
