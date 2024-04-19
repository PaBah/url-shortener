package storage

import (
	"context"
	"encoding/hex"
	"errors"
	"hash/fnv"
)

var ErrConflict = errors.New("data conflict")

type Repository interface {
	Store(ctx context.Context, URL string) (ID string, err error)
	FindByID(ctx context.Context, ID string) (URL string, err error)
	StoreBatch(ctx context.Context, URLs map[string]string) (ShortURLs map[string]string, err error)
}

func buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
