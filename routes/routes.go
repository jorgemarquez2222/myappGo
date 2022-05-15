package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name  string `json:"nombre"`
	Age   int    `json:"edad"`
	Email string `json:"email"`
}

func Server() {
	e := echo.New()

	person := &Person{
		Name:  "Jorge",
		Email: "Jorgemarquez2222@gmail.com",
		Age:   29,
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, person, "  ")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
