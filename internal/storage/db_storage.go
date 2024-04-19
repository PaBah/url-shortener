package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

type DBStorage struct {
	db *sql.DB
}

func (ds *DBStorage) initialize(ctx context.Context, databaseDSN string) (err error) {
	logger.Log().Info("DSN", zap.String("DSN", databaseDSN))
	ds.db, err = sql.Open("pgx", databaseDSN)
	if err != nil {
		logger.Log().Error("File open error", zap.Error(err))
		return
	}

	driver, err := postgres.WithInstance(ds.db, &postgres.Config{})
	if err != nil {
		logger.Log().Error("Driver open error", zap.Error(err))
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		logger.Log().Error("Migrate instance create", zap.Error(err))
		return err
	}

	err = m.Up()
	if err != nil {
		logger.Log().Error("Migration error", zap.Error(err))
	}
	return
}

func (ds *DBStorage) Store(ctx context.Context, URL string) (ID string, err error) {
	ID = buildID(URL)

	_, DBerr := ds.db.ExecContext(ctx, `INSERT INTO urls(short_url, url) VALUES ($1, $2)`, ID, URL)

	var pgErr *pgconn.PgError
	if errors.As(DBerr, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		err = ErrConflict
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

func (ds *DBStorage) FindByID(ctx context.Context, ID string) (URL string, err error) {
	row := ds.db.QueryRowContext(ctx, `SELECT url FROM urls WHERE short_url=$1`, ID)

	err = row.Scan(&URL)
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
