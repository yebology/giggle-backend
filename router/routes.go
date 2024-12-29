package router

import "github.com/gofiber/fiber/v2"

func SetUp(app *fiber.App) {

	app.Post("/api/login")
	app.Post("/api/register")

	app.Get("/api/get_post")
	app.Post("/api/create_post")
	app.Patch("/api/update_post")


}