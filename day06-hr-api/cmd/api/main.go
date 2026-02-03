package main

import (
	"context"
	"hrapi/cmd/api/routes"
	configs "hrapi/internal/config"
	"hrapi/pkg/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. set environment (bisa cmd atau system environment)
	os.Setenv("APP_ENV", "development")
	// 2. Load configuration
	config := configs.Load()

	// Initialize DB
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.CloseDB(db)

	// Run auto migration
	// if err := database.AutoMigrate(db, &models.Region{}, &models.Country{}); err != nil {
	// 	log.Printf("Warning: Auto migration failed: %v", err)
	// }

	// Set Gin mode based on environment
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Setup routes
	router := gin.Default()
	routes.SetupRoutes(router, db.DB)

	addr := config.Server.Address
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server starting on %s in %s mode", addr, config.Environment)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
