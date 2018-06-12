package main

import (
	"net/http"
	"os"

	jsonhandler "github.com/apex/log/handlers/json"
	"github.com/gorilla/pat"
	"github.com/tj/go/http/response"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func init() {
	if os.Getenv("UP_STAGE") == "" {
		log.SetHandler(text.Default)
	} else {
		log.SetHandler(jsonhandler.Default)
	}
}

func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/", trackengagement)
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func trackengagement(w http.ResponseWriter, r *http.Request) {

	// Track user, notification id, medium & url somewhere ...

	newUrl := r.URL.Query().Get("url")
	if newUrl == "" {
		response.BadRequest(w, "Missing URL")
		return
	}
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}
