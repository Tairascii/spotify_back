package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	//TODO take data from request body
	userId, err := h.manager.Auth.SignInUser("", "")

	if err != nil {
		log.Println("invalid creds")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := struct {
		Id string `json:"id"`
	}{Id: userId}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	log.Println("nice from sign up")
}
