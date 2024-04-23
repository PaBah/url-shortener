package storage

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestDBStorage_Close(t *testing.T) {
	db, mock, _ := sqlmock.New()
	dbStorage := DBStorage{db: db}
	mock.ExpectClose().WillReturnError(nil)

	err := dbStorage.Close()
	assert.NoError(t, err)
}

func TestDBStorage_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT url FROM urls WHERE short_url=$1")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"short_url"}).
			AddRow("test"))

	Data, err := ds.FindByID(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, models.NewShortURL("test"), Data, "Found message scanned correctly")
}

func TestDBStorage_Ping(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectClose().WillReturnError(nil)
	err := ds.Ping(context.Background())
	assert.NoError(t, err)
}

func TestDBStorage_Store(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls(short_url, url) VALUES ($1, $2)")).
		WithArgs("bc2c0be9", "test").WillReturnResult(sqlmock.NewResult(1, 1))

	shortURL := models.NewShortURL("test")
	_ = ds.Store(context.Background(), shortURL)
	assert.Equal(t, "bc2c0be9", shortURL.UUID, "Found message scanned correctly")
}

func TestNewDBStorage(t *testing.T) {
	_, err := NewDBStorage(context.Background(), "test")
	assert.Error(t, err, "Don't not initialize DB storage with incorrect DSN")
}

func TestDBStorage_StoreBatch(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url) VALUES($1, $2)")).
		WithArgs("bc2c0be9", "test").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url) VALUES($1, $2)")).
		WithArgs("bc2c0be9", "test").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	shortURLs := map[string]models.ShortenURL{"test1": models.NewShortURL("test"), "test2": models.NewShortURL("test")}
	err := ds.StoreBatch(context.Background(), shortURLs)
	assert.NoError(t, err, "Batch value insertion not failed")
}
