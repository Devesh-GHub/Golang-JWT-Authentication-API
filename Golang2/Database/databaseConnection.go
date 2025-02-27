package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDBURI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MongoDBURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	fmt.Println("Connected to MongoDB")
	Client = client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("hapeemail").Collection(collectionName)
}
