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

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", trackengagement)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}

func trackengagement(w http.ResponseWriter, r *http.Request) {

	qvalues := r.URL.Query()
	if len(qvalues) == 0 {
		http.Error(w, "Missing parameter values", http.StatusBadRequest)
		return
	}

	for _, input := range []string{"url", "user", "id", "medium"} {
		if qvalues.Get(input) == "" {
			http.Error(w, fmt.Sprintf("Missing %s parameter", input), http.StatusBadRequest)
			return
		}
	}

	// Track user, notification id, medium & url somewhere ...
	log.Infof("Input %v", qvalues)

	newURL := qvalues.Get("url")
	_, err := url.ParseRequestURI(newURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s is not a valid URL", newURL), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, newURL, http.StatusSeeOther)
}
