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
	"github.com/Str1kez/url-shortener/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.SetConfigFile("config/config.yaml")
	return viper.ReadInConfig()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Can't load envs from .env\n%s\n", err)
	}

	if err := initConfig(); err != nil {
		log.Fatalf("Error in parsing config file\n%s\n", err)
	}

	dbConfig := db.InitDatabaseConfig()

	database, err := dbConfig.DBConnect()
	if err != nil {
		log.Fatalf("Error in DB connection\n%s\n", err)
	}

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := urlshortener.Server{}

	dbModel := db.NewDbModel(database)
	handler := handler.Handler{Model: dbModel}

	go func() {
		host, port := viper.GetString("host"), viper.GetString("port")
		if err = server.Run(host, port, handler.InitRouters()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error on server startup\n%s\n", err)
		}
	}()

	<-terminate
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Error in shutting down\n%s\n", err)
	}
	if err = database.Close(); err != nil {
		log.Fatalf("Error in closing connection with db\n%s\n", err)
	}
}
