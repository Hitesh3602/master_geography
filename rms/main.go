package main

import (
	"log"
	"net/http"

	"github.com/Hitesh3602/master_geography/internal/config"
	"github.com/Hitesh3602/master_geography/internal/db"
	"github.com/Hitesh3602/master_geography/internal/service"
	transportHttp "github.com/Hitesh3602/master_geography/internal/transport"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	// Initialize repository and service
	geoRepo := db.NewPostgresGeographyRepository(database)
	geoService := service.NewGeographyService(geoRepo)

	// Set up HTTP transport
	handler := transportHttp.NewHTTPHandler(geoService)

	// Start the HTTP server
	log.Println("Starting server on :8083")
	if err := http.ListenAndServe(":8083", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
