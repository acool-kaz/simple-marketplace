package main

import (
	"log"
	"main/config"
	"main/internal/app"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.InitConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.NewApp(config)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
