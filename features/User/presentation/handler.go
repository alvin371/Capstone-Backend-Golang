package presentation

import (
	// Import Echo

	user "capstone/backend/features/User"
	"capstone/backend/features/User/presentation/rep"
	"capstone/backend/features/User/presentation/req"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBussiness user.Bussiness
}

func NewHandlerAccount(userBussiness user.Bussiness) *UserHandler {
	return &UserHandler{userBussiness: userBussiness}
}

func (usrHandler *UserHandler) CreateUserHandler(e echo.Context) error {
	newAccount := req.User{}

	if err := e.Bind(&newAccount); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newAccount,
	})
}

func (usrHandler *UserHandler) GetAllUserHandler(e echo.Context) error {
	result := usrHandler.userBussiness.GetAllUser("")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "All is Well",
		"data":    rep.ToUserCoreSlice(result),
	})
}

func (usrHandler *UserHandler) LoginUserHandler(e echo.Context) error {
	AccountAuth := req.UserAuth{}
	e.Bind(&AccountAuth)

	data, err := usrHandler.userBussiness.LoginUser(AccountAuth.ToUserAuth())
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserLoginResponse(data),
	})
}

func (usrHandler *UserHandler) UpdateAccountHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	// fmt.Println("Test : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := usrHandler.userBussiness.EditUser(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    rep.ToUserCore(data),
	})
}
