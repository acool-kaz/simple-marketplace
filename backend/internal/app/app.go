package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/acool-kaz/simple-marketplace/internal/config"
	httpHandler "github.com/acool-kaz/simple-marketplace/internal/delivery/http"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
	"github.com/acool-kaz/simple-marketplace/internal/service"
	"github.com/acool-kaz/simple-marketplace/pkg/postgres"
)

type app struct {
	cfg *config.Config

	db *sql.DB

	httpServer  *http.Server
	httpHandler *httpHandler.Handler
}

func InitApp(cfg *config.Config) (*app, error) {
	log.Println("init app")

	dbCfg := postgres.InitDBConfig(
		cfg.Posgres.Host,
		cfg.Posgres.Port,
		cfg.Posgres.User,
		cfg.Posgres.DbName,
		cfg.Posgres.Password,
		cfg.Posgres.SSLMode,
	)

	db, err := postgres.InitDB(dbCfg)
	if err != nil {
		return nil, err
	}

	repository := repository.InitRepository(db)
	service := service.InitService(repository)
	httpHandler := httpHandler.InitHandler(service)

	return &app{
		cfg:         cfg,
		db:          db,
		httpHandler: httpHandler,
	}, nil
}

func (a *app) RunApp() {
	log.Println("run app")
	go func() {
		err := a.startHTTP()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Println("\nReceived terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	err = a.db.Close()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("db closed")
	}
}

func (a *app) startHTTP() error {
	log.Println("start http")
	router := a.httpHandler.InitRoutes()

	a.httpServer = &http.Server{
		Handler:      router,
		Addr:         ":" + a.cfg.HttpConfig.Port,
		ReadTimeout:  time.Second * time.Duration(a.cfg.HttpConfig.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(a.cfg.HttpConfig.WriteTimeout),
	}

	return a.httpServer.ListenAndServe()
}
