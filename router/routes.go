package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
)

func SetUp(app *fiber.App) {

	app.Post("/api/login")
	app.Post("/api/register", controller.Register)

	app.Get("/api/get_post")
	app.Post("/api/create_post")
	app.Patch("/api/update_post")


}