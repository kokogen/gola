package apiserver

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}

func MakeLoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		traceId := r.Context().Value("TraceID")
		logrus.WithFields(logrus.Fields{
			"method":   r.Method,
			"url_path": r.URL.Path,
			"trace_id": traceId,
		}).Info("")
		next.ServeHTTP(w, r)
	}
}

func MakeTrackingMiddlware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "TraceID", uuid.NewString())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
