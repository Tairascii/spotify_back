package managers

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"spotify_back/repository"
)

type SongManager struct {
	repo repository.Song
}

func NewSongManager(repo repository.Song) *SongManager {
	return &SongManager{repo: repo}
}

func (man *SongManager) UploadFileSong(file *multipart.File) (string, error) {
	tempFile, err := os.CreateTemp("assets/songs", "upload-*.mp3")

	if err != nil {
		log.Printf("something went wrong while uploading song %s", err.Error())
		return "", err
	}

	defer tempFile.Close()

	fileBytes, err := io.ReadAll(*file)

	if err != nil {
		log.Printf("something went wrong while uploading song %s", err.Error())
		return "", err
	}

	tempFile.Write(fileBytes)
	return tempFile.Name(), nil
}

func (man *SongManager) CreateSong() (int, error) {
	return 0, nil
}
