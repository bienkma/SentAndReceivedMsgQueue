package middleware

import "net/http"

func CORS(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // for dev
		next.ServeHTTP(w, r)
	})
}