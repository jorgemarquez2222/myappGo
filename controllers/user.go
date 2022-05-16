package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

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

func DoSomethingChannel(c chan<- int, i int) {
	fmt.Println("Doing something numero ", i)
	time.Sleep(time.Duration(rand.Intn(i+1)) * time.Second)
	c <- i
}

func TestChannels(c echo.Context) error {
	c1 := make(chan int)

	for i := 0; i < 10; i++ {
		go DoSomethingChannel(c1, i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("")
		fmt.Println("Finished channel number", <-c1)
	}
	return c.JSON(http.StatusOK, "results")

}

func DoSomethingWg(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("Doing something numero ", i)

}

func TestWg(c echo.Context) error {
	var wg sync.WaitGroup
	len := rand.Intn(60)
	fmt.Println("cantidad a ejecutar", len)
	for i := 0; i < len; i++ {
		wg.Add(1)
		go DoSomethingWg(i, &wg)
	}
	wg.Wait()
	return c.JSON(http.StatusOK, "results")

}
