package routes

import (
	"fibr/controllers"
	"fibr/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permission", controllers.AllPermisions)

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
