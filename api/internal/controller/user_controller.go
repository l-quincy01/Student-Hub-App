package controller

import (
	"encoding/json"
	"student-hub-app/internal/model"
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userGetter  service.UserGetter
	userCreator service.UserCreator
}

func NewUserController(userGetter service.UserGetter, userCreator service.UserCreator) ControllerManager {
	return &UserController{
		userGetter, userCreator,
	}
}

func (u *UserController) Router(app *fiber.App) {
	group := app.Group("/user")
	{
		group.Get("/", u.getAllHandler)
		group.Post("/create", u.createHandler)
	}
}

func (u *UserController) getAllHandler(ctx *fiber.Ctx) error {
	users, err := u.userGetter.GetUsers(ctx.Context())
	if err != nil {
		ctx.Status(err.StatusCode())

		return err
	}

	return ctx.JSON(map[string]any{
		"data": users,
	})
}

func (u *UserController) createHandler(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()
	user := new(model.User)
	err := json.Unmarshal(reqBody, user)
	if err != nil {
		return fiber.ErrBadRequest
	}

	createErr := u.userCreator.CreateUser(ctx.Context(), *user)
	if createErr != nil {
		ctx.Status(createErr.StatusCode())

		return createErr
	}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(map[string]any{
		"message": "sucessfully created user",
	})
}
