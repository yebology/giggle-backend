package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/middleware"
	"github.com/yebology/giggle-backend/router"
)

func main() {

	app := fiber.New()

	database.ConnectDatabase()
	defer database.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(middleware.AuthMiddleware)

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))

}