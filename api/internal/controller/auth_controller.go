package controller

import (
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	googleAuthService service.GoogleAuthService
}

func NewAuthController(googleAuthService service.GoogleAuthService) AuthController {
	return AuthController{
		googleAuthService,
	}
}

func (a *AuthController) SetupRoutes(app *fiber.App) {
	group := app.Group(AuthBaseURL)
	{
		group.Get(AuthGoogleCallbackURL, a.googleCallbackHandler)
		group.Get(AuthGoogleLoginURL, a.googleLoginHandler)
	}
}

func (a *AuthController) googleCallbackHandler(ctx *fiber.Ctx) error {
	state := ctx.Query("state")

	code := ctx.Query("code")

	accessToken, err := a.googleAuthService.GetAccessToken(state, code)
	if err != nil {
		ctx.Status(err.GetStatus())

		return err
	}

	data, err := a.googleAuthService.GetUserInfo(accessToken)
	if err != nil {
		ctx.Status(err.GetStatus())

		return err
	}

	return ctx.JSON(map[string]any{
		"data": data,
	})
}

func (a *AuthController) googleLoginHandler(ctx *fiber.Ctx) error {
	url := a.googleAuthService.GetAuthCodeUrl()

	ctx.Status(fiber.StatusSeeOther)

	ctx.Redirect(url)

	return ctx.JSON(url)
}
