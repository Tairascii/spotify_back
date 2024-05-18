package handlers

import (
	"context"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) identifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(authHeader)
		if token == "" {
			return
		}

		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 {
			return
		}

		claims, err := h.manager.Auth.ParseToken(token)

		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), "userId", claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
