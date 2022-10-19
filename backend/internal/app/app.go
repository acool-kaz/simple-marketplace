package app

import (
	"fmt"
	"main/config"
	"main/internal/delivery"
	"main/internal/repository"
	"main/internal/server"
	"main/internal/service"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	db, err := repository.InitDB(a.cfg)
	if err != nil {
		return fmt.Errorf("app run: %w", err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := delivery.NewHandler(service)

	server := new(server.Server)
	if err := server.Start(a.cfg, handler.InitRoutes()); err != nil {
		return fmt.Errorf("app run: %w", err)
	}

	return nil
}
