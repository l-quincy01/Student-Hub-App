package controller

import (
	"student-hub-app/internal/model"
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type AuthController struct {
	oAuthService service.OAuthService
	userCreator  service.UserCreator
}

func NewAuthController(oAuthService service.OAuthService, userCreator service.UserCreator) ControllerManager {
	return &AuthController{
		oAuthService,
		userCreator,
	}
}

func (a *AuthController) Router(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Get("/google-login", a.googleLoginHandler)
	authGroup.Get("/google-callback", a.googleCallbackHandler)
}

func (a *AuthController) googleLoginHandler(ctx *fiber.Ctx) error {
	url := a.oAuthService.GetAuthCodeUrl()

	ctx.Status(fiber.StatusSeeOther)

	ctx.Redirect(url)

	return ctx.JSON(url)
}

func (a *AuthController) googleCallbackHandler(ctx *fiber.Ctx) error {
	state := ctx.Query("state")

	code := ctx.Query("code")

	accessToken, getTokenErr := a.oAuthService.GetAccessToken(state, code)
	if getTokenErr != nil {
		return ctx.SendStatus(getTokenErr.StatusCode())
	}

	userInfo, getUserErr := a.oAuthService.GetUserInfo(accessToken)
	if getUserErr != nil {
		ctx.SendStatus(getUserErr.StatusCode())
	}

	user := new(model.User)
	user.FirstName = userInfo.FirstName
	user.LastName = userInfo.LastName
	user.Email = userInfo.Email

	createErr := a.userCreator.CreateUser(ctx.Context(), *user)
	if createErr != nil {
		log.Errorf("unable to create user: %s, %v", user.ID, createErr)
		return ctx.SendStatus(createErr.StatusCode())
	}

	return ctx.JSON(map[string]any{
		"data": userInfo,
	})
}
