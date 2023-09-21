package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/constants"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

var RequestIDHeader = "X-Request-ID"
var prefix string
var reqID uint64

func init() {
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

	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			myID := atomic.AddUint64(&reqID, 1)
			requestID = fmt.Sprintf("%s-%09d", prefix, myID)
		}
		ctx = context.WithValue(ctx, constants.CtxKeyRequestID, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
