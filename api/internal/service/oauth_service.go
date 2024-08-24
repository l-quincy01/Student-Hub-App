package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	syserrors "student-hub-app/internal/errors"
	"student-hub-app/internal/model"

	"golang.org/x/oauth2"
)

type OAuthService struct {
	cfg *oauth2.Config
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

func (g *OAuthService) GetAccessToken(state string, code string) (string, syserrors.ApiError) {
	if state != "state" {
		return "", syserrors.UnathorizedError{Err: errors.New("state codes do not match")}
	}

	token, err := g.cfg.Exchange(context.Background(), code)
	if err != nil {
		return "", syserrors.UnathorizedError{Err: errors.New("code-token exchange failed")}
	}

	return token.AccessToken, nil
}

func (g *OAuthService) GetUserInfo(accessToken string) (model.OAuthUserInfo, syserrors.ApiError) {
	res, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken))
	if err != nil {
		return model.OAuthUserInfo{}, syserrors.BadRequestError{Err: errors.New("unable to retrieve user data")}
	}

	if res.StatusCode != http.StatusOK {
		return model.OAuthUserInfo{}, syserrors.InternalServerError{Err: errors.New("unable to process request to retrieve user data")}
	}

	userInfo := new(model.OAuthUserInfo)
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		return model.OAuthUserInfo{}, syserrors.InternalServerError{Err: fmt.Errorf("unable to unmarshal user data from response: %v", err)}
	}

	return *userInfo, nil
}
