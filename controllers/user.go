package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name  string `json:"nombre"`
	Age   int    `json:"edad"`
	Email string `json:"email"`
}

func User(c echo.Context) error {
	person := &Person{
		Name:  "Jorge",
		Email: "Jorgemarquez2222@gmail.com",
		Age:   39,
	}
	return c.JSON(http.StatusOK, person)
}
