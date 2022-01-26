package routes

import (
	"capstone/backend/config"
	"capstone/backend/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	jwt := e.Group("")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderAccept},
	}))
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	// News CRUD
	e.GET("/news", presenter.NewsPresentation.GetAllNewsHandler)
	e.GET("/news/:id", presenter.NewsPresentation.GetNewsByIDHandler)
	e.POST("/news/create", presenter.NewsPresentation.CreateNewsHandler)
	e.PATCH("/news/update/:id", presenter.NewsPresentation.EditNewsHandler)

	// User Credential
	e.POST("/user/register", presenter.UserPresentation.CreateUserHandler)
	e.POST("/user/login", presenter.UserPresentation.LoginUserHandler)
	jwt.GET("/user", presenter.UserPresentation.GetAllUserHandler)
	jwt.GET("/user/:id", presenter.UserPresentation.GetUserById)
	e.PUT("/user/:id", presenter.UserPresentation.UpdateAccountHandler)

	// Online Class CRUD
	e.GET("/online-class", presenter.OnlineClassPresentation.GetAllClassHandler)
	e.GET("/online-class/:id", presenter.OnlineClassPresentation.GetClassByIdHandler)
	jwt.POST("/online-class/create", presenter.OnlineClassPresentation.CreateClassHandler)
	jwt.PATCH("/online-class/edit/:id", presenter.OnlineClassPresentation.UpdateClassHandler)
	jwt.PATCH("/online-class/delete/:id", presenter.OnlineClassPresentation.DeleteClassHandler)

	// Offline Class CRUD
	e.GET("/offline-class", presenter.PresenterOfflineClassPresentation.GetAllOfflineClassHandler)
	e.GET("/offline-class/:id", presenter.PresenterOfflineClassPresentation.GetOfflineClassByIdHandler)
	e.POST("/offline-class/create", presenter.PresenterOfflineClassPresentation.CreateOfflineClassHandler)
	e.PATCH("/offline-class/edit/:id", presenter.PresenterOfflineClassPresentation.UpdateOfflineClassHandler)
	e.PATCH("/offline-class/delete/:id", presenter.PresenterOfflineClassPresentation.DeleteOfflineClassHandler)

	// Booking Offline
	jwt.GET("/booking-offline", presenter.BookingOfflinePresentation.GetListBookingOffline)
	e.POST("/booking-offline/create", presenter.BookingOfflinePresentation.InsertMemberBookingOffline)

	// Booking Online

	jwt.GET("/booking-online", presenter.BookingOnlinePresentation.GetListBookingOnline)
	e.POST("/booking-online/create", presenter.BookingOnlinePresentation.InsertMemberBookingOnline)
	return e
}
