package main

import (
	"capstone/backend/driver"
	_middleware "capstone/backend/middlewares"
	"capstone/backend/migrate"
	"capstone/backend/routes"
)

func main() {
	driver.InitDB()
	migrate.AutoMigrate()
	e := routes.New()
	// Log Middleware
	_middleware.LogMiddlewareInit(e)

	e.Start(":8000")

}
