package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"student-hub-app/internal/model"

	"golang.org/x/oauth2"
)

type OAuthService struct {
	cfg *oauth2.Config
}

type OAuthUserInfo struct {
	// TODO: populate with user info payload fields from Google
}

func NewOAuthService(cfg *oauth2.Config) OAuthService {
	return OAuthService{
		cfg,
	}
}

func (g *OAuthService) GetAuthCodeUrl() string {
	url := g.cfg.AuthCodeURL("state")

	return url
}

func (g *OAuthService) GetAccessToken(state string, code string) (string, model.ApiError) {
	if state != "state" {
		return "", model.UnathorizedError{Err: errors.New("state codes do not match")}
	}

	token, err := g.cfg.Exchange(context.Background(), code)
	if err != nil {
		return "", model.UnathorizedError{Err: errors.New("code-token exchange failed")}
	}

	return token.AccessToken, nil
}

func (g *OAuthService) GetUserInfo(accessToken string) (any, model.ApiError) {
	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken))
	if err != nil {
		return "", model.BadRequestError{Err: errors.New("unable to retrieve user data")}
	}

	if res.StatusCode != http.StatusOK {
		return "", model.InternalServerError{Err: errors.New("unable to process request to retrieve user data")}
	}

	userInfo := new(OAuthUserInfo)
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		return "", model.InternalServerError{Err: fmt.Errorf("unable to unmarshal user data from response: %v", err)}
	}

	return userInfo, nil
}
