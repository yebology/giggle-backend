package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/oauth"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
)

func SetUp(app *fiber.App) {

	// done check
	app.Get("/oauth/google", oauth.GoogleAuth)
	app.Get("/oauth/redirect", oauth.GoogleRedirect)

	// done check postman
	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)

	// done check postman
	app.Get("/api/get_posts", controller.GetPosts)
	app.Post("/api/create_post", middleware.UserMiddleware, controller.CreatePost)
	app.Patch("/api/update_post/:id", middleware.UserMiddleware, middleware.PostOwnerMiddleware, controller.UpdatePost)
	app.Delete("/api/delete_post/:id", middleware.UserMiddleware, middleware.PostOwnerMiddleware, controller.DeletePost)

	// done check postman
	app.Get("/api/get_groups/:user_id", middleware.UserMiddleware, controller.GetUserGroups)
	app.Post("/api/create_group", middleware.UserMiddleware, controller.CreateGroup)
	app.Patch("/api/invite_member_to_group/:id", middleware.UserMiddleware, middleware.GroupOwnerMiddleware, controller.InviteMember)

}