package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"spotify_back/models"
	"spotify_back/pkg"
)

type signInUser struct {
	Email    string
	Password string
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var userInput signInUser

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		pkg.JSONResponse(w, map[string]string{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	token, err := h.manager.Auth.SignInUser(userInput.Email, userInput.Password)

	if err != nil {
		log.Println("invalid credentials")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := struct {
		AccessToken string `json:"accessToken"`
	}{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		pkg.JSONResponse(w, map[string]string{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	userId, err := h.manager.Auth.SignUpUser(userInput)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pkg.JSONResponse(w, map[string]int{"id": userId}, http.StatusOK)
	return
}
