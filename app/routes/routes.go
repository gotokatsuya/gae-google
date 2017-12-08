package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/gotokatsuya/gae-google/app/controllers"
)

var (
	// App is
	App = negroni.New(negroni.NewRecovery())
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.IndexHandler).Methods(http.MethodGet)

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Path("/login").HandlerFunc(controllers.AuthLoginHandler).Methods(http.MethodGet)
	authRouter.Path("/google").HandlerFunc(controllers.AuthGoogleHandler).Methods(http.MethodGet)
	authRouter.Path("/google/callback").HandlerFunc(controllers.AuthGoogleCallbackHandler).Methods(http.MethodGet)

	App.UseHandler(router)
}
