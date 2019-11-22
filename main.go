package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rwenz3l/lmdb/server"
)

func main() {
	// CLI Flags
	var dataDir = flag.String("dataDir", "static/", "Directory for Storing Blob Data")
	var httpPort = flag.Int("port", 8080, "Port for the HTTP WebServer")
	flag.Parse()

	// Create Dependencies
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	_, err := sql.Open("postgres", "user=lmdb dbname=media sslmode=verify-full")
	if err != nil {
		level.Error(logger).Log("Could not connect to postgres")
		panic(err)
	}
	level.Debug(logger).Log("msg", "Connected to PostgreSQL")

	// Router + Server
	r := mux.NewRouter()
	srv := server.New(logger)

	// Routes
	// Static FileServer
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*dataDir))))

	// Home
	r.HandleFunc("/", srv.HomeHandler).Methods("GET")

	// API - Search
	r.HandleFunc("/search/tv", srv.TVSearchHandler).Methods("GET")
	r.HandleFunc("/search/movie", srv.MovieSearchHandler).Methods("GET")
	r.HandleFunc("/search/anime", srv.AnimeSearchHandler).Methods("GET")
	r.HandleFunc("/search/person", srv.PersonSearchHandler).Methods("GET")

	// API - TV [GET, POST, PUT, DELETE]
	// API - Movie [GET, POST, PUT, DELETE]
	// API - Anime  [GET, POST, PUT, DELETE]

	// Start WebServer
	level.Debug(logger).Log("msg", fmt.Sprintf("Starting Webserver on %d", *httpPort))
	err = http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), r)
	if err != nil {
		level.Debug(logger).Log("msg", "Error on httpServer", "err", err)
	}
}
