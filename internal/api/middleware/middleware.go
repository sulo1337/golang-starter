package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

var RequestIDHeader = "X-Request-ID"

type Middleware interface {
	RequestID(next http.Handler) http.Handler
	Authenticated(next http.Handler) http.Handler
	Logger(next http.Handler) http.Handler
}

type middleware struct {
	logger      *slog.Logger
	reqIDPrefix string
	reqID       uint64
}

func NewMiddleware(logger *slog.Logger) Middleware {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	return &middleware{
		logger:      logger,
		reqIDPrefix: fmt.Sprintf("%s/%s", hostname, b64[0:10]),
	}
}
