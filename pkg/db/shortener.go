package db

import (
	"time"

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

func (s *ShortenerPostgres) Get(shortUrl string) (string, error) {
	fullUrl := "https://ya.ru"

	return fullUrl, nil
}

func (s *ShortenerPostgres) GetInfo(secret string) (*schema.InfoResponse, error) {
	shortUrl := "FHSFL"
	longUrl := "https://ya.ru"
	numClicks := 32
	createdAt := time.Now().Format(time.RFC3339)

	return &schema.InfoResponse{
		ShortUrl:       shortUrl,
		LongUrl:        longUrl,
		NumberOfClicks: numClicks,
		CreatedAt:      createdAt,
	}, nil
}

func (s *ShortenerPostgres) Delete(secret string) error {

	return nil
}
