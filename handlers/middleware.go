package handlers

import (
	"context"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) identifyUser(r *http.Request) {
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

	ctx := r.Context()
	ctx = context.WithValue(ctx, "userId", claims.UserId)
}
