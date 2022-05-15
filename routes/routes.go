package routes

import (
	controllers "github.com/jorgemarquez2222/myappGo/controllers"
	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()
	e.GET("/", controllers.User)
	e.Logger.Fatal(e.Start(":1323"))
}
