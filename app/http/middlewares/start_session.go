package middlewares

import (
	"net/http"

	"github.com/yangliang4488/goblog/pkg/session"
)

func StartSession(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session.StartSession(rw, r)

		next.ServeHTTP(rw, r)
	})
}
