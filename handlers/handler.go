package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	r.Use(middleware.Logger)
	r.Mount("/auth", authHandlers(h))
	r.Mount("/playlist", playlistHandlers(h))
	r.Mount("/song", songHandlers(h))
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

func songHandlers(h *Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
			h.uploadSong(w, r)
		})
	})

	return rg
}

func playlistHandlers(h *Handler) http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			h.createPlaylist(w, r)
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			h.editPlaylist(w, r)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			h.getPlaylist(w, r)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			h.deletePlaylist(w, r)
		})
		r.Post("/{id}/add-song/{songId}", func(w http.ResponseWriter, r *http.Request) {
			h.addSongToPlaylist(w, r)
		})
		r.Delete("/{id}/delete-song/{songId}", func(w http.ResponseWriter, r *http.Request) {
			h.deleteSongToPlaylist(w, r)
		})
	})
	return rg
}
