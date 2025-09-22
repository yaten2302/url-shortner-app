package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

// initialise MongoDB connection
func Connect() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGO_DB")
	if uri == "" || dbName == "" {
		log.Fatal("MONGO_URI or MONGO_DB not set in environment")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	Client = client
	DB = client.Database(dbName)

	fmt.Println("Connected to MongoDB:", dbName)
	return DB
}

// Disconnect the MongoDB connection
func Disconnect() {
	if Client != nil {
		if err := Client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}
}
