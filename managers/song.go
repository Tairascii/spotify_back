package managers

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"spotify_back/models"
	"spotify_back/repository"
	"time"
)

const (
	defaultCoverPath = "assets/thumbnails/default_cover.png"
)

type SongManager struct {
	repo repository.Song
}

func NewSongManager(repo repository.Song) *SongManager {
	return &SongManager{repo: repo}
}

func (man *SongManager) UploadSong(file *multipart.File, title, author string, userId int) (int, error) {
	tempFile, err := os.CreateTemp("assets/songs", "upload-*.mp3")

	if err != nil {
		log.Printf("something went wrong while uploading song %s", err.Error())
		return 0, err
	}

	defer tempFile.Close()

	fileBytes, err := io.ReadAll(*file)

	if err != nil {
		log.Printf("something went wrong while uploading song %s", err.Error())
		return 0, err
	}

	tempFile.Write(fileBytes)

	id, err := man.repo.CreateSong(models.Song{
		Author:    author,
		Title:     title,
		SongPath:  tempFile.Name(),
		ImagePath: defaultCoverPath,
		UserId:    userId,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Printf("something went wrong while creating song %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (man *SongManager) CreateSong() (int, error) {
	return 0, nil
}
