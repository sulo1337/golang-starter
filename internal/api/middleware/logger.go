package middleware

import (
	"fmt"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"time"
)

func (m middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		reqID := getReqID(r.Context())
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		ww := chiMiddleware.NewWrapResponseWriter(w, r.ProtoMajor)
		m.logger.LogAttrs(
			r.Context(),
			slog.LevelDebug,
			"request started",
			slog.String("req_id", reqID),
			slog.String("http_scheme", scheme),
			slog.String("http_proto", r.Proto),
			slog.String("http_method", r.Method),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
			slog.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)),
		)
		defer func() {
			m.logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"request complete",
				slog.String("req_id", reqID),
				slog.Int("resp_status", ww.Status()),
				slog.Int("resp_byte_length", ww.BytesWritten()),
				slog.Float64("resp_elapsed_ms", float64(time.Since(t1).Nanoseconds())/1000000.0),
			)
		}()
		next.ServeHTTP(ww, r.WithContext(r.Context()))
	})
}
