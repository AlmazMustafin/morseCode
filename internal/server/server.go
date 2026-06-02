package server

import (
	"log"
	"net/http"
	"time"
 	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.ServeIndexHandler)
	mux.Handle("/upload", handlers.UploadHandler(logger))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}
}

func (s *Server) Start() error {
	s.Logger.Printf("Starting server on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}
