package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/rwenz3l/lmdb/server"
	"github.com/rwenz3l/lmdb/server/anime"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// CLI Flags
	var dataDir = flag.String("dataDir", "static/", "Directory for Storing Blob Data")
	var httpPort = flag.Int("port", 8080, "Port for the HTTP WebServer")
	flag.Parse()

	// Logger
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))

	// MongoDB
	// docker run -p 27017:27017 mongo:latest
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("lmdb")

	// Router
	router := mux.NewRouter()
	// Services
	anime := anime.NewService(logger, db.Collection("anime"))

	// Server
	srv := server.New(logger, router, anime)
	srv.InitRoutes(*dataDir)

	// Run the Server
	srv.Run(*httpPort)
}
