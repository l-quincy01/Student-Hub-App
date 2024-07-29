package controller

import (
	"student-hub-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	oauthService service.OAuthService
}

func NewOAuthController(googleAuthService service.OAuthService) AuthController {
	return AuthController{
		googleAuthService,
	}
}

const (
	baseURL               = "/auth"
	AuthGoogleCallbackURL = "/google_callback"
	AuthGoogleLoginURL    = "/google_login"
)

func (a *AuthController) Router(app *fiber.App) {
	group := app.Group(baseURL)
	{
		group.Get(AuthGoogleLoginURL, a.googleLoginHandler)
		group.Get(AuthGoogleCallbackURL, a.googleCallbackHandler)
	}
}

func (a *AuthController) googleLoginHandler(ctx *fiber.Ctx) error {
	url := a.oauthService.GetAuthCodeUrl()

	ctx.Status(fiber.StatusSeeOther)

	ctx.Redirect(url)

	return ctx.JSON(url)
}

func (a *AuthController) googleCallbackHandler(ctx *fiber.Ctx) error {
	state := ctx.Query("state")

	code := ctx.Query("code")

	accessToken, err := a.oauthService.GetAccessToken(state, code)
	if err != nil {
		ctx.Status(err.StatusCode())

		return err
	}

	data, err := a.oauthService.GetUserInfo(accessToken)
	if err != nil {
		ctx.Status(err.StatusCode())

		return err
	}

	return ctx.JSON(map[string]any{
		"data": data,
	})
}
