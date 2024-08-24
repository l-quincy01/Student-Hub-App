package controller

import "github.com/gofiber/fiber/v2"

type ControllerManager interface {
	Router(app *fiber.App)
}