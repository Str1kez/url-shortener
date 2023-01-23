package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/Str1kez/url-shortener/schema"
	"github.com/jmoiron/sqlx"
)

type ShortenerPostgres struct {
	db *sqlx.DB
}

const (
	tablename    = "urlshortener"
	shortUrlSize = 5
)

func NewShortenerPostgres(db *sqlx.DB) *ShortenerPostgres {
	return &ShortenerPostgres{db}
}

func (s *ShortenerPostgres) Create(url string) (*schema.ShortenerResponse, error) {

	var shortUrl, secretKey string
	var err error

	query := fmt.Sprintf(`SELECT short_url, id FROM %s WHERE long_url=$1`, tablename)
	row := s.db.QueryRow(query, url)
	if err = row.Scan(&shortUrl, &secretKey); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == nil {
		return &schema.ShortenerResponse{
			ShortUrl:  shortUrl,
			SecretKey: secretKey,
		}, nil
	}

	shortCh := make(chan string)
	errCh := make(chan error)

	go makeShort(s.db, shortCh, errCh)

	select {
	case err = <-errCh:
		return nil, err
	case shortUrl = <-shortCh:
		query = fmt.Sprintf(`INSERT INTO %s (short_url, long_url) VALUES ($1, $2) RETURNING id`, tablename)
		row = s.db.QueryRow(query, shortUrl, url)
		if err = row.Scan(&secretKey); err != nil {
			return nil, err
		}
		return &schema.ShortenerResponse{
			ShortUrl:  shortUrl,
			SecretKey: secretKey,
		}, nil
	case <-time.After(time.Second):
		return nil, &Timeout{}
	}
}

func (s *ShortenerPostgres) Get(shortUrl string) (string, error) {
	var longUrl string

	query := fmt.Sprintf(`SELECT long_url FROM %s WHERE short_url=$1`, tablename)
	row := s.db.QueryRow(query, shortUrl)
	if err := row.Scan(&longUrl); err != nil {
		if err == sql.ErrNoRows {
			return "", &NoResultFound{}
		}
		return "", err
	}
	return longUrl, nil
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

func makeShort(db *sqlx.DB, cancelCh chan<- string, errCh chan<- error) {
	var isExists bool

	pattern := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456890"
	rand.Seed(time.Now().UnixNano())
	for {
		patternInRune := []rune(pattern)
		rand.Shuffle(len(patternInRune), func(i, j int) {
			patternInRune[i], patternInRune[j] = patternInRune[j], patternInRune[i]
		})
		shortUrl := string(patternInRune[:shortUrlSize])
		query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE short_url=$1)`, tablename)
		row := db.QueryRow(query, shortUrl)
		if err := row.Scan(&isExists); err != nil {
			errCh <- err
			return
		}
		if !isExists {
			cancelCh <- shortUrl
			return
		}
	}
}
