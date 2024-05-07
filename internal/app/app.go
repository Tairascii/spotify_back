package app

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"spotify_back/repository"
	"time"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) initConfigs(configName string) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}

func (app *App) Start(configName string) {
	if err := app.initConfigs(configName); err != nil {
		log.Fatalf("something went wrong initing config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("something went wrong loading env %s", err.Error())
	}
	db, conErr := repository.NewPostgresDB(repository.Config{
		DBName:   viper.GetString("db.name"),
		Host:     viper.GetString("db.host"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
	})

	if conErr != nil {
		log.Fatalf("Something went wrong while connecting to database %s", conErr.Error())
	}

	repository := repository.NewRepository(db)
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
