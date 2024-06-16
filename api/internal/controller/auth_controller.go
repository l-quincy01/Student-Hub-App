package controller

import (
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

const (
	GoogleCallbackUrl = "/google_callback"
	GoogleLogin       = "/google_login"
)

type AuthController struct {
	googleAuthService service.GoogleAuthService
}

func NewAuthController(googleAuthService service.GoogleAuthService) AuthController {
	return AuthController{
		googleAuthService,
	}
}

func (a *AuthController) Route(app *fiber.App) {
	app.Get(GoogleLogin, a.googleLoginHandler)
	app.Get(GoogleCallbackUrl, a.googleCallbackHandler)
}

func (a *AuthController) googleCallbackHandler(ctx *fiber.Ctx) error {
	state := ctx.Query("state")

	code := ctx.Query("code")

	token, err := a.googleAuthService.TokenExchange(state, code)
	if err != nil {
		ctx.Status(500)

		return err
	}

	return ctx.JSON(map[string]any{
		"access_token": token,
	})
}

func (a *AuthController) googleLoginHandler(ctx *fiber.Ctx) error {
	url := a.googleAuthService.ConsentRedirectUrl()

	ctx.Status(fiber.StatusSeeOther)

	ctx.Redirect(url)

	return ctx.JSON(url)
}
