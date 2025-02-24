package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go-net_http-project/docs"
	"go-net_http-project/internal/config"
	"go-net_http-project/internal/helpers"
	"go-net_http-project/internal/middleware"
	"go-net_http-project/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title		Api Documentation
// @version		1.0
// @description	Api docs for net-http template
// @host		localhost:3001
// @BasePath	/
func main() {
	if err := config.LoadConfig("./internal/config/config.json"); err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	opts := helpers.DefaultOptions()
	opts.MinLength = 12

	if err := helpers.Initialize(opts); err != nil {
		log.Fatalf("Failed to initialize sqids encoder: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbPool, err := config.InitializeDatabase(ctx)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer dbPool.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		cancel()
		dbPool.Close()
	}()

	router := http.NewServeMux()
	routes.RegisterRoutes(router, dbPool)

	router.Handle("/docs/", httpSwagger.WrapHandler)

	middlewareStack := middleware.CreateStack(middleware.Logging)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")),
		Handler: middlewareStack(router),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :3001: %v\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
