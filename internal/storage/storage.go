package storage

import (
	"context"
	"encoding/hex"
	"hash/fnv"
)

type Repository interface {
	Store(ctx context.Context, Data string) (ID string, duplicate bool)
	FindByID(ctx context.Context, ID string) (Data string, err error)
	StoreBatch(ctx context.Context, URLs map[string]string) (ShortURLs map[string]string, err error)
}

func buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
