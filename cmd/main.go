package main

import (
	"log"

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
}
