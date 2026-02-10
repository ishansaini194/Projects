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

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file")
	}

	MongoDb := os.Getenv("MONGO_URL")
	if MongoDb == "" {
		log.Fatal("MONGO_URL not found in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenConnection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}

// func OpenConnection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
// 	return collection
// }
