package controller

import (
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

const (
	baseUrl = "/users"
)

func (u UserController) Router(app *fiber.App) {
	// group := app.Group(baseURL)
	// {

	// }
}
