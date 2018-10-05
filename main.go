package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

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
	if newURL == "" {
		http.Error(w, "url searchParam is empty", http.StatusBadRequest)
		return
	}

	u, err := url.ParseRequestURI(newURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid URL", newURL), http.StatusBadRequest)
		return
	}

	// We only want to redirect to unee-t.com sites
	uneetDomain := "unee-t.com" // 2 last parts
	splitHost := strings.Split(u.Host, ".")
	if len(splitHost) < 2 {
		log.Errorf("Host %s, length %d", u.Host, len(splitHost))
		http.Error(w, fmt.Sprintf("%s is not a valid unee-t.com URL", newURL), http.StatusBadRequest)
		return
	}
	if strings.Join(splitHost[len(splitHost)-2:], ".") != uneetDomain {
		http.Error(w, fmt.Sprintf("%s is not a valid unee-t.com URL", newURL), http.StatusBadRequest)
		return
	}

	for _, input := range []string{"email", "id"} {
		if qvalues.Get(input) == "" {
			http.Error(w, fmt.Sprintf("Missing %s parameter", input), http.StatusBadRequest)
			return
		}
	}

	log.WithFields(log.Fields{
		"id":    qvalues.Get("id"),
		"email": qvalues.Get("email"),
		"url":   qvalues.Get("url"),
	}).Info("input")

	http.Redirect(w, r, newURL, http.StatusSeeOther)
}

func fail(w http.ResponseWriter, r *http.Request) {
	log.Warn("5xx")
	http.Error(w, "5xx", http.StatusInternalServerError)
}
