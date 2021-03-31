package user

import (
	"app/main/backend_examples/mvc/models/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserHandler() *fiber.App {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		username := make([]byte, len(c.FormValue("username")))
		copy(username, c.FormValue("username"))
		if found := user.CheckUsernameExist(string(username)); found {
			return c.Status(fiber.StatusConflict).SendString("User already exists")
		}
		user.CreateUser(string(username))
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/create", func(c *fiber.Ctx) error {
		return c.Render("user/create", fiber.Map{}, "layouts/main")
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		userId := make([]byte, len(c.Params("id")))
		copy(userId, c.Params("id"))
		if u := user.FindUserByID(string(userId)); u != nil {
			return c.Render("user/index", fiber.Map{
				"ID":   string(userId),
				"Name": u.Name,
				"Data": u.Data,
			}, "layouts/main")
		}
		return c.Redirect("/error")
	})

	app.Patch("/:id", func(c *fiber.Ctx) error {
		userId := make([]byte, len(c.Params("id")))
		copy(userId, c.Params("id"))
		data, _ := strconv.Atoi(c.FormValue("data"))
		if user.UpdateUserData(string(userId), data) {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		userId := make([]byte, len(c.Params("id")))
		copy(userId, c.Params("id"))
		user.DeleteUser(string(userId))
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}
