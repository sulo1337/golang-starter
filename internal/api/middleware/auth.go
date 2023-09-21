package middleware

import (
	"net/http"
)

func (m middleware) Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Debug("middleware::Authenticated")
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
