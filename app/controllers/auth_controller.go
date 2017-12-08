package controllers

import (
	"net/http"

	"google.golang.org/appengine"

	"github.com/gotokatsuya/gae-google/lib/template"

	"github.com/gotokatsuya/gae-google/app/services/auth"
)

// AuthLoginHandler ...
func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	template.Render.HTML(w, http.StatusOK, "auth/login", nil)
}

func AuthGoogleHandler(w http.ResponseWriter, r *http.Request) {
	svc := auth.NewService()
	state := svc.GenerateState()
	url := svc.GoogleAuthURL(state)
	svc.SaveState(state, r, w)
	http.Redirect(w, r, url, http.StatusFound)
}

func AuthGoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	svc := auth.NewService()
	if r.URL.Query().Get("state") != svc.State(r) {
		http.Error(w, "invalid state", http.StatusUnauthorized)
		return
	}
	if err := svc.GoogleLogin(ctx, r.URL.Query().Get("code")); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	http.Redirect(w, r, "https://eure.jp", http.StatusFound)
}
