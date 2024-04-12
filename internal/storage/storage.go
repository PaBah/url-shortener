package storage

import (
	"context"
	"encoding/hex"
	"hash/fnv"
)

type Repository interface {
	Store(ctx context.Context, Data string) (ID string)
	FindByID(ctx context.Context, ID string) (Data string, err error)
}

func buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
