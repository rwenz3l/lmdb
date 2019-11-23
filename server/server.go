package server

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/rwenz3l/lmdb/server/anime"
)

// Server contains all dependencies for the API Server
type Server struct {
	logger log.Logger
	router *mux.Router
	anime  *anime.Service
}

// New provides a new Instance of a Server
func New(logger log.Logger, router *mux.Router, anime *anime.Service) *Server {
	server := &Server{
		logger: logger,
		router: router,
		anime:  anime,
	}
	return server
}

// InitRoutes sets all endpoints and handlers
func (s *Server) InitRoutes(dataDir string) {

	// Routes
	// Static FileServer
	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dataDir))))

	// Home
	s.router.HandleFunc("/", s.HomeHandler).Methods("GET")

	// API - Search
	s.router.HandleFunc("/search/anime", s.anime.SearchHandler).Methods("GET")
	// s.router.HandleFunc("/search/movie", s.movie.SearchHandler).Methods("GET")
	// s.router.HandleFunc("/search/tv", s.tv.SearchHandler).Methods("GET")

	// API - Anime  [GET, POST, PUT, DELETE]
	s.router.HandleFunc("/anime", s.anime.CreateHandler).Methods("POST")

	// API - TV [GET, POST, PUT, DELETE]
	// API - Movie [GET, POST, PUT, DELETE]
}

// Run starts the WebServer and listens until panic or SIGKILL
func (s *Server) Run(port int) {
	// Start WebServer
	level.Debug(s.logger).Log("msg", fmt.Sprintf("Starting Webserver on %d", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
	if err != nil {
		level.Debug(s.logger).Log("msg", "Error on httpServer", "err", err)
	}
}

// HomeHandler for the API
func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(s.logger).Log("msg", "Called HomeHandler")
	fmt.Fprintf(w, "Welcome to LMDB!")
}
