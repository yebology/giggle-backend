package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
	"github.com/yebology/giggle-backend/middleware/http"
	"github.com/yebology/giggle-backend/oauth"
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
	app.Post("/api/create_post", middleware.AuthMiddleware, controller.CreatePost)
	app.Patch("/api/update_post/:id", middleware.AuthMiddleware, http.PostOwnerMiddleware, controller.UpdatePost)
	app.Delete("/api/delete_post/:id", middleware.AuthMiddleware, http.PostOwnerMiddleware, controller.DeletePost)

	// done check postman
	app.Get("/api/get_groups/:user_id", middleware.AuthMiddleware, controller.GetUserGroups)
	app.Post("/api/create_group", middleware.AuthMiddleware, controller.CreateGroup)
	app.Patch("/api/invite_member_to_group/:id", middleware.AuthMiddleware, http.GroupOwnerMiddleware, controller.InviteMember)

}