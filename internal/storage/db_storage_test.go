package storage

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
	assert.Equal(t, "test", Data, "Found message scanned correctly")
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
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls(short_url, url) VALUES ($1, $2) ON CONFLICT DO NOTHING")).
		WithArgs("test", "test").WillReturnResult(sqlmock.NewResult(1, 1))

	ID, _ := ds.Store(context.Background(), "test")
	assert.Equal(t, "bc2c0be9", ID, "Found message scanned correctly")
}

func TestNewDBStorage(t *testing.T) {
	_, err := NewDBStorage(context.Background(), "test")
	assert.Error(t, err, "Don't not initialize DB storage with incorrect DSN")
}

func TestDBStorage_migrate(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectExec(regexp.QuoteMeta(`
		CREATE TABLE IF NOT EXISTS urls (
		    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		    short_url VARCHAR NOT NULL UNIQUE,
		    url VARCHAR NOT NULL UNIQUE
		    )
		`)).WillReturnResult(sqlmock.NewResult(1, 1))

	err := ds.migrate(context.Background())
	assert.NoError(t, err, "Don't not initialize DB storage with incorrect DSN")
}

func TestDBStorage_StoreBatch(t *testing.T) {
	db, mock, _ := sqlmock.New()
	ds := &DBStorage{
		db: db,
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url) VALUES($1, $2) ON CONFLICT DO NOTHING")).
		WithArgs("bc2c0be9", "test").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO urls (short_url, url) VALUES($1, $2) ON CONFLICT DO NOTHING")).
		WithArgs("bc2c0be9", "test").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	shortURLs, err := ds.StoreBatch(context.Background(), map[string]string{"test1": "test", "test2": "test"})
	assert.NoError(t, err, "Batch value insertion not failed")
	assert.Equal(t, map[string]string{"test1": "bc2c0be9", "test2": "bc2c0be9"}, shortURLs, "All batch stored")
}
