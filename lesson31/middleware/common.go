package middleware

import (
	"log"
	"net/http"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
