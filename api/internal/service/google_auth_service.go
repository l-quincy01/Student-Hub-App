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

type GoogleAuthService struct {
	conf *oauth2.Config
}

func NewGoogleAuthService(conf *oauth2.Config) GoogleAuthService {
	return GoogleAuthService{
		conf,
	}
}

func (g *GoogleAuthService) ConsentRedirectUrl() string {
	url := g.conf.AuthCodeURL("state")

	return url
}

func (g *GoogleAuthService) TokenExchange(state string, code string) (string, error) {
	if state != "state" {
		return "", model.UnathorizedError{Err: errors.New("state codes do not match")}
	}

	token, err := g.conf.Exchange(context.Background(), code)
	if err != nil {
		return "", model.UnathorizedError{Err: errors.New("code-token exchange failed")}
	}

	return token.AccessToken, nil
}

func (g *GoogleAuthService) GetUserData(token string) (string, error) {
	userInfoUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token)

	resp, err := http.Get(userInfoUrl)
	if err != nil {
		return "", model.BadRequestError{Err: errors.New("unable to retrieve user data")}
	}

	if resp.StatusCode != http.StatusOK {
		return "", model.InternalServerError{Err: errors.New("unable to process request to retrieve user data")}
	}

	var userData string
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		return "", model.InternalServerError{Err: fmt.Errorf("unable to unmarshal user data from response. err: %v", err)}
	}

	return userData, nil
}
