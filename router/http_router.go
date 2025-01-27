package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/controller"
	"github.com/yebology/giggle-backend/middleware"
	"github.com/yebology/giggle-backend/middleware/http"
	"github.com/yebology/giggle-backend/oauth"
)

// SetUp initializes the application routes and associates them with their respective controllers.
// It sets up routes for OAuth authentication, user registration and login, posts management, and group management.
func SetUp(app *fiber.App) {

	// Route for Google OAuth authentication. The user will be redirected to Google for login.
	// Once the user authenticates, they are redirected back to our app.
	app.Get("/oauth/google", oauth.GoogleAuth)  // Initiates OAuth with Google
	app.Get("/oauth/redirect", oauth.GoogleRedirect)  // Redirect URI for OAuth callback

	// User authentication routes for login and registration
	// Done check with Postman
	app.Post("/api/login", controller.Login)  // Handles user login
	app.Post("/api/register", controller.Register)  // Handles user registration

	// Post management routes
	// Done check with Postman
	app.Get("/api/get_posts", controller.GetPosts)  // Retrieves all posts
	app.Post("/api/create_post", middleware.AuthMiddleware, controller.CreatePost)  // Creates a new post (requires authentication)
	app.Patch("/api/update_post/:id", middleware.AuthMiddleware, http.PostOwnerMiddleware, controller.UpdatePost)  // Updates an existing post (requires authentication and ownership check)
	app.Delete("/api/delete_post/:id", middleware.AuthMiddleware, http.PostOwnerMiddleware, controller.DeletePost)  // Deletes a post (requires authentication and ownership check)

	app.Get("/api/get_proposals/:user_id", middleware.AuthMiddleware, controller.GetProposals)
	app.Post("/api/create_proposal", middleware.AuthMiddleware, controller.CreateProposal)
	app.Patch("/api/accept_proposal/:id", middleware.AuthMiddleware, http.BuyerMiddleware, controller.AcceptProposal)

	// Group management routes
	// Done check with Postman
	app.Get("/api/get_groups/:user_id", middleware.AuthMiddleware, controller.GetUserGroups)  // Retrieves groups owned by a user (requires authentication)
	app.Post("/api/create_group", middleware.AuthMiddleware, controller.CreateGroup)  // Creates a new group (requires authentication)
	app.Patch("/api/invite_member_to_group/:id", middleware.AuthMiddleware, http.GroupOwnerMiddleware, controller.InviteMember)  // Invites a member to a group (requires authentication and ownership check)
	
}
