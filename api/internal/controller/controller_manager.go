package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Router(app *fiber.App)
}

