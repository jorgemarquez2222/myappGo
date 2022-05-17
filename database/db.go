package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Manager interface {
	GetCollection(nameCollection string) *mongo.Collection
}

type manager struct {
	client *mongo.Client
}

var Mgr Manager

func ConnDB() {
	fmt.Println("Connecting to MongoDB...")
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
	Mgr = &manager{client: client}
}

// GetCollection implements Manager
func (m *manager) GetCollection(nameCollection string) *mongo.Collection {
	collection := m.client.Database("test").Collection(nameCollection)
	return collection
}
