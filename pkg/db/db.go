package db

import (
	"github.com/Str1kez/url-shortener/schema"
	"github.com/jmoiron/sqlx"
)

type Shortener interface {
	Create(url string) (*schema.ShortenerResponse, error)
	Get(shortUrl string)
	GetInfo(secret string)
	Delete(secret string)
}

type DbModel struct {
	Shortener
}

func NewDbModel(db *sqlx.DB) *DbModel {
	return &DbModel{
		Shortener: NewShortenerPostgres(db),
	}
}
