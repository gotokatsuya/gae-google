package auth

import (
	"context"
	"net/http"

	"github.com/gotokatsuya/gae-google/lib/auth/google"
	"github.com/gotokatsuya/gae-google/lib/random"
	"github.com/gotokatsuya/gae-google/lib/sessions"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateState() string {
	return random.UUID()
}

func (s *Service) GoogleAuthURL(state string) string {
	auth := google.NewAuth()
	return auth.CodeURL(state)
}

func (s *Service) SaveState(state string, r *http.Request, w http.ResponseWriter) {
	session, _ := sessions.Get(r, "google")
	session.Values["state"] = state
	sessions.Save(r, w, session)
}

func (s *Service) State(r *http.Request) string {
	session, _ := sessions.Get(r, "google")
	state := session.Values["state"]
	if state == nil {
		return ""
	}
	return state.(string)
}

func (s *Service) GoogleLogin(ctx context.Context, code string) error {
	auth := google.NewAuth()
	_, err := auth.Login(ctx, code)
	if err != nil {
		return err
	}
	return nil
}
