package google

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

func newConfig() *oauth2.Config {
	const (
		authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
		tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
	)
	return &oauth2.Config{
		ClientID:     os.Getenv("AUTH_GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH_GOOGLE_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"email", "profile"},
		RedirectURL: os.Getenv("AUTH_GOOGLE_REDIRECT_URL"),
	}
}

type Auth struct {
	config *oauth2.Config
}

func NewAuth() *Auth {
	return &Auth{
		config: newConfig(),
	}
}

func (a Auth) CodeURL(state string) string {
	return a.config.AuthCodeURL(state)
}

type User struct {
	ID    string
	Email string
}

func (a Auth) Login(ctx context.Context, code string) (*User, error) {
	token, err := a.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if !token.Valid() {
		return nil, errors.New("invaild token")
	}
	service, err := v2.New(a.config.Client(ctx, token))
	if err != nil {
		return nil, err
	}
	tokenInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return &User{ID: tokenInfo.UserId, Email: tokenInfo.Email}, nil
}
