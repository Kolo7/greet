package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/handler"
)

type UserAgentMiddleware struct {
}

func NewUserAgentMiddleware() *UserAgentMiddleware {
	return &UserAgentMiddleware{}
}

func (m *UserAgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		handler.BreakerHandler(r.Method, r.URL.Path, nil)(next).ServeHTTP(w, r)
	}
}
