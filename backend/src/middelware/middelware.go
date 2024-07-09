package middelware

import (
	"log"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("| %s\t| %s\n", r.Method, r.URL)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}
