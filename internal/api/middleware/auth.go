package middleware

import (
	"fmt"
	"net/http"
)

func (m middleware) Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware::MustBeAuthenticated called")
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
