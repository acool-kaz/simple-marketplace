package main

import (
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/app"
	"github.com/acool-kaz/simple-marketplace/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.InitConfig()

	app, err := app.InitApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.RunApp()
}
