package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	jsonhandler "github.com/apex/log/handlers/json"

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

func muxEngine() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/fail", fail)
	mux.HandleFunc("/", trackengagement)
	return mux
}

func main() {
	addr := ":" + os.Getenv("PORT")
	if err := http.ListenAndServe(addr, muxEngine()); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func trackengagement(w http.ResponseWriter, r *http.Request) {

	qvalues := r.URL.Query()
	if len(qvalues) == 0 {
		http.Error(w, "Missing parameter values", http.StatusBadRequest)
		return
	}

	newURL := qvalues.Get("url")
	_, err := url.ParseRequestURI(newURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid URL", newURL), http.StatusBadRequest)
		return
	}

	for _, input := range []string{"user", "id", "medium"} {
		if qvalues.Get(input) == "" {
			http.Error(w, fmt.Sprintf("Missing %s parameter", input), http.StatusBadRequest)
			return
		}
	}

	log.WithFields(log.Fields{
		"id":     qvalues.Get("id"),
		"user":   qvalues.Get("user"),
		"url":    qvalues.Get("url"),
		"medium": qvalues.Get("medium"),
	}).Info("input")

	http.Redirect(w, r, newURL, http.StatusSeeOther)
}

func fail(w http.ResponseWriter, r *http.Request) {
	log.Warn("5xx")
	http.Error(w, "5xx", http.StatusInternalServerError)
}
