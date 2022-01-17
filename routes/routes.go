package routes

import (
	"capstone/backend/config"
	"capstone/backend/factory"

	"github.com/labstack/echo"
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

	// News
	e.GET("/news", presenter.NewsPresentation.GetAllNewsHandler)
	return e
}
