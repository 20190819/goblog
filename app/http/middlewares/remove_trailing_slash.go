package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

// 除首页以外，移除所有请求路径后面的斜杠

func RemoveTrailingSlash(next http.Handler) http.Handler {
	fmt.Println("hello test")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(rw, r)
	})
}
