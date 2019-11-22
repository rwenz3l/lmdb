package server

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/rwenz3l/lmdb/media"
)

// Server contains all dependencies for the API Server
type Server struct {
	// Logger, Repository
	logger log.Logger
	media  media.Repository
}

// New provides a new Instance of a Server
func New(logger log.Logger) *Server {
	server := &Server{
		logger: logger,
	}
	return server
}

// HomeHandler returns the default homepage message
func (h *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(h.logger).Log("msg", "HomeHandler")
	fmt.Fprintf(w, "Welcome to the LMDB API!")
}

// TVSearchHandler searches for a given tv-show within the database
func (h *Server) TVSearchHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(h.logger).Log("msg", "TVSearchHandler")
	fmt.Fprintf(w, "Search Handler for TV Shows")
}

// MovieSearchHandler searches for a given movie within the database
func (h *Server) MovieSearchHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(h.logger).Log("msg", "MovieSearchHandler")
	fmt.Fprintf(w, "Search Handler for Movies")
}

// AnimeSearchHandler searches for a given anime within the database
func (h *Server) AnimeSearchHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(h.logger).Log("msg", "AnimeSearchHandler")
	fmt.Fprintf(w, "Search Handler for Animes")
}

// PersonSearchHandler searches for a given person within the database
func (h *Server) PersonSearchHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(h.logger).Log("msg", "PersonSearchHandler")
	fmt.Fprintf(w, "Search Handler for Animes")
}
