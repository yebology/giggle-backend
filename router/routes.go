package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
)

func SetUp(app *fiber.App) {

	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)
	app.Post("/api/check_account", controller.CheckAccount)

	app.Get("/api/get_post", controller.GetPost)
	app.Post("/api/create_post", controller.CreatePost)
	app.Patch("/api/update_post/:id", controller.UpdatePost)

}