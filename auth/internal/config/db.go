package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL is not defined")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Print("Error with connecting to Mongo!")
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Print("Cannot ping MongoDB!")
		log.Fatal(err)
	}

	log.Print("Successfully connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectMongo()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Cluster0").Collection(collectionName)
	return collection
}
