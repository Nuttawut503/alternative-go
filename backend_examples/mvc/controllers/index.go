package controllers

import (
	"app/main/backend_examples/mvc/models/user"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		users := user.GetAllUsers()
		if len(users) == 0 {
			users = nil
		}
		return c.Render("index", fiber.Map{
			"Users": users,
		}, "layouts/main")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return c.Render("error", fiber.Map{
			"Message": "Page Not Found",
		}, "layouts/main")
	})

	return app
}
