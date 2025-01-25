package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/handler"
	"github.com/yebology/giggle-backend/router"
)

func main() {

	// Initialize the Fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDatabase()
	
	// Ensure that the database connection is closed once the main function finishes execution
	defer database.DisconnectDatabase()

	// Set up CORS (Cross-Origin Resource Sharing) to allow requests from any origin
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowHeaders: "*", // Alow all headers
	}))

	// Set up HTTP routes using the router
	router.SetUp(app)

	// Set up WebSocket connections using the handler
	handler.SetUp(app)

	// Start the application on port 8080
	log.Fatal(app.Listen(":8080"))

}