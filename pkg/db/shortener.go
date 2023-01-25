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

	query := fmt.Sprintf(`SELECT short_url, id FROM %s WHERE long_url=$1 AND active=true`, tablename)
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

	query := fmt.Sprintf(`SELECT long_url FROM %s WHERE short_url=$1 AND active=true`, tablename)
	row := s.db.QueryRowx(query, shortUrl)
	if err := row.Scan(&longUrl); err != nil {
		if err == sql.ErrNoRows {
			return "", &NoResultFound{}
		}
		return "", err
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return "", err
	}

	query = fmt.Sprintf(`UPDATE %s
                       SET number_of_clicks = number_of_clicks + 1
                       WHERE short_url=$1 AND active=true`, tablename)
	row = tx.QueryRowx(query, shortUrl)

	if err := row.Err(); err != nil {
    tx.Rollback()
		return "", err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}

	return longUrl, nil
}

func (s *ShortenerPostgres) GetInfo(secret string) (*schema.InfoResponse, error) {
	var response schema.InfoResponse

	query := fmt.Sprintf(`SELECT short_url AS ShortUrl, long_url AS LongUrl, number_of_clicks AS NumberOfClicks,
                        dt_created AS CreatedAt
                        FROM %s WHERE id=$1 AND active=true`, tablename)

	if err := s.db.Get(&response, query, secret); err != nil {
		if err == sql.ErrNoRows {
			return nil, &NoResultFound{}
		}
		return nil, err
	}

	return &response, nil
}

func (s *ShortenerPostgres) Delete(secret string) error {

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`UPDATE %s
                        SET active=false
                        WHERE id=$1`, tablename)
	row := tx.QueryRowx(query, secret)
	if err = row.Err(); err != nil {
    tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

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
		query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE short_url=$1 AND active=true)`, tablename)
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
