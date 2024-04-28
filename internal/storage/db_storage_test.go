package storage

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
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
	mock.ExpectQuery(regexp.QuoteMeta("SELECT url, user_id, is_deleted FROM urls WHERE short_url=$1")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"short_url", "user_id", "is_deleted"}).
			AddRow("test", 1, false))

	Data, err := ds.FindByID(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, models.ShortenURL{UUID: "test", OriginalURL: "test", UserID: "1", DeletedFlag: false}, Data, "Found message scanned correctly")
}

func TestDBStorage_FindByID_with_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT url FROM urls WHERE short_url=$1")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"short_url"}).
			AddRow(nil))

	_, err := ds.FindByID(context.Background(), "test")
	assert.Error(t, err)
	//assert.Equal(t, models.NewShortURL("test"), Data, "Found message scanned correctly")
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

	shortURL := models.NewShortURL("test", "1")
	_ = ds.Store(context.Background(), shortURL)
	assert.Equal(t, "bc2c0be9", shortURL.UUID, "Found message scanned correctly")
}

func TestDBStorage_Store_with_error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls(short_url, url, user_id) VALUES ($1, $2, $3)")).
		WithArgs("bc2c0be9", "test", "1").WillReturnError(&pgconn.PgError{Code: "23505"})

	shortURL := models.NewShortURL("test", "1")
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, 1)
	err := ds.Store(ctx, shortURL)
	assert.Error(t, err, "duplicate")
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
	mock.ExpectQuery(regexp.QuoteMeta("SELECT short_url FROM urls")).
		WillReturnRows(sqlmock.NewRows([]string{"short_url"}).
			AddRow("test"))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url, user_id) VALUES($1, $2, $3)")).
		WithArgs("bc2c0be9", "test", "1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	shortURLs := map[string]models.ShortenURL{"test1": models.NewShortURL("test", "1")}
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, 1)
	err := ds.StoreBatch(ctx, shortURLs)
	assert.NoError(t, err, "Batch value insertion not failed")
}

func TestDBStorage_StoreBatch_with_error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT short_url FROM urls")).
		WillReturnError(&pgconn.PgError{Code: "777"})

	shortURLs := map[string]models.ShortenURL{"test1": models.NewShortURL("test", "1")}
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, 1)
	err := ds.StoreBatch(ctx, shortURLs)
	assert.Error(t, err, "Batch value insertion failed")
}

func TestDBStorage_StoreBatch_parse_error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT short_url FROM urls")).
		WillReturnRows(sqlmock.NewRows([]string{"short_url"}).
			AddRow("test"))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url) VALUES($1, $2)")).
		WithArgs("bc2c0be9", "test").WillReturnError(&pgconn.PgError{Code: "777"})
	mock.ExpectCommit()

	shortURLs := map[string]models.ShortenURL{"test1": models.NewShortURL("test", "1")}
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, 1)
	err := ds.StoreBatch(ctx, shortURLs)
	assert.Error(t, err, "Batch value insertion failed")
}

func TestDBStorage_GetAllUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT url, short_url, user_id FROM urls WHERE user_id=$1")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"url", "short_url", "user_id"}).
			AddRow("url", "test", "test"))
	ctx := context.WithValue(context.Background(), auth.ContextUserKey, "test")
	Data, err := ds.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.Equal(t, []models.ShortenURL{models.ShortenURL{UUID: "test", OriginalURL: "url", UserID: "test"}}, Data, "Found message scanned correctly")
}
