package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	log.Println("nice")
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	log.Println("nice from sign up")
}
