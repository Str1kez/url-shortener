package db

import (
	"github.com/Str1kez/url-shortener/schema"
	"github.com/jmoiron/sqlx"
)

type ShortenerPostgres struct {
	db *sqlx.DB
}

func NewShortenerPostgres(db *sqlx.DB) *ShortenerPostgres {
	return &ShortenerPostgres{db}
}

func (s *ShortenerPostgres) Create(url string) (*schema.ShortenerResponse, error) {
	// logic
	shortUrl, secretKey := "something", "idk"

	return &schema.ShortenerResponse{
		ShortUrl:  shortUrl,
		SecretKey: secretKey,
	}, nil
}

func (s *ShortenerPostgres) Get(shortUrl string) {

}

func (s *ShortenerPostgres) GetInfo(secret string) {

}

func (s *ShortenerPostgres) Delete(secret string) {

}
