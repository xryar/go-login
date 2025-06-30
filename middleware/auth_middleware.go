package middleware

import (
	"login-app/helper"
	"login-app/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		writeUnAuthorized(w)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)
	userId, err := helper.ValidateJWT(token)
	if err != nil {
		writeUnAuthorized(w)
		return
	}

	ctx := r.Context()
	ctx = helper.ContextWithUserId(ctx, userId)

	middleware.Handler.ServeHTTP(w, r.WithContext(ctx))
}

func writeUnAuthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}

	helper.WriteToResponseBody(w, webResponse)
}
