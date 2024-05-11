package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"spotify_back/managers"
)

type Handler struct {
	manager *managers.Manager
}

func NewHandler(manager *managers.Manager) *Handler {
	return &Handler{manager: manager}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/auth", authHandlers(h))

	return r
}

func authHandlers(h *Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
			h.signIn(w, r)
		})
		r.Post("/sign-up", func(w http.ResponseWriter, r *http.Request) {
			h.signUp(w, r)
		})
		r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
			h.refresh(w, r)
		})
	})

	return rg
}
