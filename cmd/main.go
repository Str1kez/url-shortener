package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	urlshortener "github.com/Str1kez/url-shortener"
	"github.com/Str1kez/url-shortener/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Can't load envs from .env\n%s\n", err)
	}

	dbConfig := db.InitDatabaseConfig()

	db, err := dbConfig.DBConnect()
	if err != nil {
		log.Fatalf("Error in DB connection\n%s\n", err)
	}

	db.Ping()

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := urlshortener.Server{}
	go func() {
		if err = server.Run("8001"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error on server startup\n%s\n", err)
		}
	}()

	<-terminate
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Error in shutting down\n%s\n", err)
	}
}
