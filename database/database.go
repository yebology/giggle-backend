package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func GetDatabase() *mongo.Database {

	if Client == nil {
		log.Fatalf("MongoDB is not initialized!")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env!")
	}

	DB_NAME := os.Getenv("DB_NAME")

	return  Client.Database(DB_NAME)

}