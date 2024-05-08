package main

import "spotify_back/internal/app"

func main() {
	a := app.NewApp()
	a.Start("config")
}
