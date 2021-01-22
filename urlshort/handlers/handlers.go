package handlers

import (
	"net/http"
)

//MapHandler is a function
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if redirectionURL, found := pathsToUrls[r.URL.Path]; found {
			http.Redirect(w, r, redirectionURL, 301)
			return
		}
		fallback.ServeHTTP(w, r)
	}

}
