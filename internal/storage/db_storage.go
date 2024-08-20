package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/PaBah/url-shortener.git/db"
	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
)

// DBStorage - model of Repository storage on top of Data Base
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

// Store - stores shortened URL in DB
func (ds *DBStorage) Store(ctx context.Context, shortURL models.ShortenURL) (err error) {
	_, DBerr := ds.db.ExecContext(ctx,
		`INSERT INTO urls(short_url, url, user_id) VALUES ($1, $2, $3)`, shortURL.UUID, shortURL.OriginalURL, shortURL.UserID)

	var pgErr *pgconn.PgError
	if errors.As(DBerr, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		err = ErrConflict
	}

	return
}

// StoreBatch - stores batch of shortened URLs in DB
func (ds *DBStorage) StoreBatch(ctx context.Context, shortURLsMap map[string]models.ShortenURL) (err error) {
	rows, err := ds.db.QueryContext(ctx, `SELECT short_url FROM urls`)
	if err != nil {
		return err
	}
	defer rows.Close()

	shortURLs := make([]string, 0)
	var shortURL string
	for rows.Next() {
		err = rows.Scan(&shortURL)
		if err != nil {
			return err
		}
		shortURLs = append(shortURLs, shortURL)
	}
	err = rows.Err()

	tx, err := ds.db.Begin()
	if err != nil {
		return err
	}
	for _, shortURL := range shortURLsMap {
		if !slices.Contains(shortURLs, shortURL.UUID) {
			shortURLs = append(shortURLs, shortURL.UUID)
			_, err = tx.ExecContext(ctx,
				"INSERT INTO urls (short_url, url, user_id) VALUES($1, $2, $3)", shortURL.UUID, shortURL.OriginalURL, shortURL.UserID)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				_ = tx.Rollback()
				return
			}
		}
	}
	return tx.Commit()
}

// FindByID - filter and returns shortened URL by short ID
func (ds *DBStorage) FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error) {
	row := ds.db.QueryRowContext(ctx, `SELECT url, user_id, is_deleted FROM urls WHERE short_url=$1`, ID)
	var URL string
	var userID string
	var deletedFlag bool
	err = row.Scan(&URL, &userID, &deletedFlag)

	if err != nil {
		return
	}

	shortURL = models.ShortenURL{OriginalURL: URL, UUID: ID, UserID: userID, DeletedFlag: deletedFlag}
	return
}

// GetAllUsers - returns all shortened URLs of the User from context
func (ds *DBStorage) GetAllUsers(ctx context.Context) (shortURLs []models.ShortenURL, err error) {
	var rows *sql.Rows
	rows, err = ds.db.QueryContext(ctx, `SELECT url, short_url, user_id FROM urls WHERE user_id=$1`, ctx.Value(auth.ContextUserKey).(string))
	if err != nil {
		return
	}
	err = rows.Err()
	defer rows.Close()

	shortURLs = make([]models.ShortenURL, 0)
	for rows.Next() {
		var shortURL models.ShortenURL
		err = rows.Scan(&shortURL.OriginalURL, &shortURL.UUID, &shortURL.UserID)
		if err != nil {
			return nil, err
		}
		shortURLs = append(shortURLs, shortURL)
	}
	return
}

// AsyncCheckURLsUserID - async checking if URL belongs to the User from context
func (ds *DBStorage) AsyncCheckURLsUserID(userID string, shortURLCh chan string) chan string {
	addRes := make(chan string)
	go func() {
		defer close(addRes)

		for data := range shortURLCh {

			shortURL, err := ds.FindByID(context.Background(), data)
			var result string
			fmt.Println("shortURL", shortURL)
			fmt.Println("err", err)
			if err == nil && shortURL.UserID == userID {
				result = shortURL.UUID
			}

			addRes <- result
		}
	}()
	return addRes
}

// DeleteShortURLs - delete shortened URLs from Data Base
func (ds *DBStorage) DeleteShortURLs(ctx context.Context, shortURLs []string) (err error) {
	_, err = ds.db.ExecContext(ctx, `UPDATE urls SET is_deleted = TRUE WHERE urls.short_url = ANY($1)`, pq.Array(shortURLs))

	return
}

// GetStats - return amount of users and amount of url in the system
func (ds *DBStorage) GetStats(ctx context.Context) (urls int, users int, err error) {
	err = ds.db.QueryRowContext(ctx, `SELECT COUNT(id), COUNT(DISTINCT user_id) FROM urls`).Scan(&urls, &users)

	return
}

// Ping - check if connection to Data Base is fine
func (ds *DBStorage) Ping(ctx context.Context) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	return ds.db.PingContext(ctxWithTimeout)
}

// Close - close connection to Data Base
func (ds *DBStorage) Close() error {
	return ds.db.Close()
}

// NewDBStorage - create instance of DBStorage
func NewDBStorage(ctx context.Context, databaseDSN string) (DBStorage, error) {
	store := DBStorage{}
	err := store.initialize(ctx, databaseDSN)
	return store, err
}
