package storage

import (
	"context"
	"database/sql"
	"errors"
	"slices"
	"time"

	"github.com/PaBah/url-shortener.git/db"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBStorage struct {
	db *sql.DB
}

func (ds *DBStorage) initialize(ctx context.Context, databaseDSN string) (err error) {

	ds.db, err = sql.Open("pgx", databaseDSN)
	if err != nil {
		return
	}

	driver, err := iofs.New(db.MigrationsFS, "migrations")
	if err != nil {
		return err
	}

	d, err := postgres.WithInstance(ds.db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", driver, "psql_db", d)
	if err != nil {
		return err
	}

	_ = m.Up()
	return
}

func (ds *DBStorage) Store(ctx context.Context, shortURL models.ShortenURL) (err error) {
	_, DBerr := ds.db.ExecContext(ctx,
		`INSERT INTO urls(short_url, url) VALUES ($1, $2)`, shortURL.UUID, shortURL.OriginalURL)

	var pgErr *pgconn.PgError
	if errors.As(DBerr, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		err = ErrConflict
	}

	return
}

func (ds *DBStorage) StoreBatch(ctx context.Context, shortURLsMap map[string]models.ShortenURL) (err error) {
	rows, _ := ds.db.QueryContext(ctx, `SELECT short_url FROM urls`)
	defer rows.Close()
	shortURLs := make([]string, 0)
	var shortURL string
	for rows.Next() {
		_ = rows.Scan(&shortURL)
		shortURLs = append(shortURLs, shortURL)
	}

	tx, err := ds.db.Begin()
	if err != nil {
		return err
	}
	for _, shortURL := range shortURLsMap {
		if !slices.Contains(shortURLs, shortURL.UUID) {
			shortURLs = append(shortURLs, shortURL.UUID)
			_, err = tx.ExecContext(ctx,
				"INSERT INTO urls (short_url, url) VALUES($1, $2)", shortURL.UUID, shortURL.OriginalURL)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				_ = tx.Rollback()
				return
			}
		}
	}
	return tx.Commit()
}

func (ds *DBStorage) FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error) {
	row := ds.db.QueryRowContext(ctx, `SELECT url FROM urls WHERE short_url=$1`, ID)
	var URL string
	err = row.Scan(&URL)

	if err != nil {
		return
	}

	shortURL = models.NewShortURL(URL)
	return
}

func (ds *DBStorage) Ping(ctx context.Context) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	return ds.db.PingContext(ctxWithTimeout)
}

func (ds *DBStorage) Close() error {
	return ds.db.Close()
}

func NewDBStorage(ctx context.Context, databaseDSN string) (DBStorage, error) {
	store := DBStorage{}
	err := store.initialize(ctx, databaseDSN)
	return store, err
}
