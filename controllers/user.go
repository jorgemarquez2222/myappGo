package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name  string `json:"nombre"`
	Age   int    `json:"edad"`
	Email string `json:"email"`
}

type PlaceHolder struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func User(c echo.Context) error {
	person := &Person{
		Name:  "Jorge",
		Email: "Jorgemarquez2222@gmail.com",
		Age:   39,
	}
	return c.JSON(http.StatusOK, person)
}

func TestRquest(c echo.Context) error {
	placeholder := []PlaceHolder{}
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &placeholder)
	return c.JSON(http.StatusOK, placeholder)
}
