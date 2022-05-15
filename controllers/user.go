package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client = ConnDB()
var users = client.Database("test").Collection("users")

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

func ConnDB() *mongo.Client {
	godotenv.Load()
	var uri = os.Getenv("url_base")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

type PersonMongo struct {
	Name string `json:"nombre"`
	Rut  string `json:"rut"`
}

func TestRquest(c echo.Context) error {
	placeholder := []PlaceHolder{}
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &placeholder)
	return c.JSON(http.StatusOK, placeholder)
}

func TestMongo(c echo.Context) error {
	cursor, err := users.Find(c.Request().Context(),
		bson.D{{Key: "name", Value: "Jorge"}})
	if err != nil {
		panic(err)
	}
	var results []PersonMongo
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, results)
}
