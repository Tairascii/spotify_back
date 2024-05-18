package handlers

import (
	"log"
	"net/http"
	"spotify_back/pkg"
)

func (h *Handler) uploadSong(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("song")

	if err != nil {
		log.Printf("something went wrong while uploading song %s", err.Error())
		pkg.JSONResponse(w, map[string]string{"message": err.Error()}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	title := r.FormValue("title")
	author := r.FormValue("author")
	userId, ok := pkg.GetContextValue(r.Context(), "userId").(int)

	if !ok {
		pkg.JSONResponse(w, map[string]string{"message": "invalid user id"}, http.StatusBadRequest)
		return
	}

	id, err := h.manager.Song.UploadSong(&file, title, author, userId)

	if err != nil {
		pkg.JSONResponse(w, map[string]string{"message": err.Error()}, http.StatusBadRequest)
	}

	pkg.JSONResponse(w, map[string]int{"id": id}, http.StatusOK)
	return
}
