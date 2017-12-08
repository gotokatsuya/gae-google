package app

import (
	"net/http"

	"github.com/gotokatsuya/gae-google/app/routes"
)

func init() {
	http.Handle("/", routes.App)
}
