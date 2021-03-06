package routes

import (
	controllers "github.com/jorgemarquez2222/myappGo/controllers"
	db "github.com/jorgemarquez2222/myappGo/database"

	"github.com/labstack/echo/v4"
)

func Server() {
	db.ConnDB()
	e := echo.New()
	e.GET("/", controllers.User)
	e.GET("/test", controllers.TestRquest)
	e.GET("/testMongo", controllers.TestMongo)
	e.GET("/testChannels", controllers.TestChannels)
	e.GET("/testWg", controllers.TestWg)

	e.Logger.Fatal(e.Start(":1323"))
}
