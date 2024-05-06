package app

import (
	"log"
	"net/http"
	"time"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Start() {
	db, conErr := "later", false
	if conErr {
		log.Fatal("Failed to connect to database")
	}
	repos := "new repos with passed db"
	service := "new service with passed repos"
	handler := "new handlers with passed service"
	srv := &http.Server{
		Addr:         ":" + "8080", // TODO create config file
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      handler, // TODO change to normal handler
	}

	srv.ListenAndServe()
}
