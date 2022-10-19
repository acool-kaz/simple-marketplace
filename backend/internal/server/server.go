package server

import (
	"fmt"
	"main/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(cfg *config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
	fmt.Printf("Server start: http://localhost:%d/\n", cfg.Port)
	return s.httpServer.ListenAndServe()
}
