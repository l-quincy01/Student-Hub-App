package config

import (
	"context"
	"fmt"
	"os"
	"student-hub-app/internal/controller"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	googleClientId     = "GOOGLE_CLIENT_ID"
	googleClientSecret = "GOOGLE_CLIENT_SECRET"
)

func SetupGoogleOauth(ctx context.Context) *oauth2.Config {
	redirectUrl := fmt.Sprintf("%s%s", os.Getenv("BASE_URL"), controller.AuthGoogleCallbackURL)

	cfg := &oauth2.Config{
		ClientID:     os.Getenv(googleClientId),
		ClientSecret: os.Getenv(googleClientSecret),
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return cfg
}
