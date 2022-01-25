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

	e.GET("/online-class", presenter.OnlineClassPresentation.GetAllClassHandler)
	e.GET("/online-class/:id", presenter.OnlineClassPresentation.GetClassByIdHandler)
	jwt.POST("/online-class/create", presenter.OnlineClassPresentation.CreateClassHandler)
	jwt.PATCH("/online-class/edit/:id", presenter.OnlineClassPresentation.UpdateClassHandler)
	jwt.PATCH("/online-class/delete/:id", presenter.OnlineClassPresentation.DeleteClassHandler)
	// User Credential
	e.POST("/user/register", presenter.UserPresentation.CreateUserHandler)
	e.POST("/user/login", presenter.UserPresentation.LoginUserHandler)
	jwt.GET("/user", presenter.UserPresentation.GetAllUserHandler)
	jwt.GET("/user/:id", presenter.UserPresentation.GetUserById)
	e.PUT("/user/:id", presenter.UserPresentation.UpdateAccountHandler)

	return e
}
