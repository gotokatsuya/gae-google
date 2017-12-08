package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("something-very-secret"))
)

func New(name string) *sessions.Session {
	return sessions.NewSession(store, name)
}

func Get(r *http.Request, name string) (*sessions.Session, error) {
	return store.Get(r, name)
}

func Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return store.Save(r, w, session)
}
