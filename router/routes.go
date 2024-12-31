package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
)

func SetUp(app *fiber.App) {

	// done check postman
	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)
	app.Post("/api/check_account", controller.CheckAccount)

	// 
	app.Get("/api/get_post", controller.GetPost)
	app.Post("/api/create_post", middleware.UserMiddleware, controller.CreatePost)
	app.Patch("/api/update_post/:id", middleware.UserMiddleware, middleware.PostOwnerMiddleware, controller.UpdatePost)
	app.Delete("/api/delete_post/:id", middleware.UserMiddleware, middleware.PostOwnerMiddleware, controller.DeletePost)

	// 
	app.Get("/api/get_group/:user_id", middleware.UserMiddleware, controller.GetUserGroups)
	app.Post("/api/create_group", middleware.UserMiddleware, controller.CreateGroup)
	app.Patch("/api/invite_member_to_group/:id", middleware.UserMiddleware, controller.InviteMemberToGroup)

}