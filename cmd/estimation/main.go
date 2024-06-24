package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/celsopires1999/estimation/configs"
	"github.com/jackc/pgx/v5/pgxpool"

	httpHandler "github.com/celsopires1999/estimation/internal/infra/http"
)

func main() {
	ctx := context.Background()

	configs := configs.LoadConfig(".", "")
	dbpool, err := pgxpool.New(ctx, configs.DBConn)

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	v1 := httpHandler.Handler(ctx, dbpool)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.Port),
		Handler: v1,
	}

	// Channel to listen for operating system signals
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Received interrupt signal, starting graceful shutdown
		log.Println("Received interrupt signal, starting graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error during graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("HTTP server running on port", configs.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error starting HTTP server: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("HTTP server finished")
}
