package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variable to hold the MongoDB client instance
var Client *mongo.Client

// GetDatabase retrieves the MongoDB database instance using the global client.
// Ensures the client is initialized and fetches the database name from environment variables.
func GetDatabase() *mongo.Database {

	// Check if the MongoDB client is initialized
	if Client == nil {
		log.Fatalf("MongoDB is not initialized!")
	}

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env: %s", err)
	}

	// Retrieve the database name from the environment variables
	DB_NAME := os.Getenv("DB_NAME")

	// Return the database instance using the provided database name
	return Client.Database(DB_NAME)
}

// ConnectDatabase initializes a connection to MongoDB and assigns it to the global client variable.
// It also validates the connection by pinging the database.
func ConnectDatabase() {

	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load from .env: %s", err)
	}

	// Retrieve the MongoDB URI from the environment variables
	MONGO_URI := os.Getenv("MONGO_URI")

	// Set MongoDB client options with the provided URI
	clientOption := options.Client().ApplyURI(MONGO_URI)

	// Create a new MongoDB client and connect to the database
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatalf("Error while connecting to MongoDB: %s", err)
	}
	Client = client

	// Test the connection by pinging the MongoDB server
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while pinging MongoDB: %s", err)
	}

	// Log success message if connection is successful
	log.Println("Successfully connected to MongoDB!")
}

// DisconnectDatabase safely disconnects the MongoDB client and logs the result.
func DisconnectDatabase() {

	// Check if the MongoDB client is initialized
	if Client == nil {
		log.Fatalf("MongoDB is not initialized!")
	}

	// Disconnect the client from the MongoDB server
	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting from MongoDB: %s", err)
	}

	// Log success message if disconnection is successful
	log.Println("Successfully disconnected from MongoDB!")
}
