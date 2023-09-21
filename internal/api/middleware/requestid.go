package middleware

import (
	"context"
	"fmt"
	"github.com/sulo1337/cleanarch-go/internal/constants"
	"net/http"
	"sync/atomic"
)

func (m middleware) RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			myID := atomic.AddUint64(&m.reqID, 1)
			requestID = fmt.Sprintf("%s-%09d", m.reqIDPrefix, myID)
		}
		ctx = context.WithValue(ctx, constants.CtxKeyRequestID, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func getReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(constants.CtxKeyRequestID).(string); ok {
		return reqID
	}
	return ""
}
