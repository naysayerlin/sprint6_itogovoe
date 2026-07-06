package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewServ(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	serv := &Server{
		logger:     logger,
		httpServer: httpServer,
	}
	return serv
}

func (server *Server) Start() error {
	server.logger.Printf("Server is running on port %s", server.httpServer.Addr)
	return server.httpServer.ListenAndServe()
}

func (server *Server) GetHTTPServer() *http.Server {
	return server.httpServer
}
