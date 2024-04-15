package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBStorage struct {
	db *sql.DB
}

func (ds *DBStorage) init(ctx context.Context, databaseDSN string) (err error) {
	ds.db, err = sql.Open("pgx", databaseDSN)
	if err != nil {
		return
	}
	err = ds.migrate(ctx)
	return
}

func (ds *DBStorage) migrate(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	_, err = ds.db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS urls (
		    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		    short_url VARCHAR NOT NULL UNIQUE,
		    url VARCHAR NOT NULL UNIQUE
		    )
		`)
	return
}

func (ds *DBStorage) Store(ctx context.Context, Data string) (ID string, duplicate bool) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ID = buildID(Data)
	const PG_DUPLICATE_KEY_ERROR_CODE string = "23505"
	_, err := ds.db.ExecContext(ctx, `INSERT INTO urls(short_url, url) VALUES ($1, $2)`, ID, Data)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == PG_DUPLICATE_KEY_ERROR_CODE {
		duplicate = true
	}

	return
}

func (ds *DBStorage) StoreBatch(ctx context.Context, URLs map[string]string) (ShortURLs map[string]string, err error) {
	// начинаем транзакцию
	tx, err := ds.db.Begin()
	if err != nil {
		return nil, err
	}
	ShortURLs = make(map[string]string)
	for k, v := range URLs {
		ID := buildID(v)
		_, err := tx.ExecContext(ctx,
			"INSERT INTO urls (short_url, url) VALUES($1, $2) ON CONFLICT DO NOTHING", ID, v)
		ShortURLs[k] = ID
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}
	return ShortURLs, tx.Commit()
}

func (ds *DBStorage) FindByID(ctx context.Context, ID string) (Data string, err error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	row := ds.db.QueryRowContext(ctx, `SELECT url FROM urls WHERE short_url=$1`, ID)

	err = row.Scan(&Data)
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
	err := store.init(ctx, databaseDSN)
	return store, err
}
