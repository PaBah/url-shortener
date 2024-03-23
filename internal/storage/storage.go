package storage

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

type Repository interface {
	Store(string) string
	FindByID(string) (string, error)
}

type InMemoryStorage struct {
	state map[string]string
}

func (cs *InMemoryStorage) Store(Data string) (ID string) {
	if cs.state == nil {
		cs.state = make(map[string]string)
	}

	ID = cs.buildID(Data)
	cs.state[ID] = Data
	return
}

func (cs *InMemoryStorage) FindByID(ID string) (Data string, err error) {
	if cs.state == nil {
		cs.state = make(map[string]string)
	}

	Data, found := cs.state[ID]
	if !found {
		return Data, fmt.Errorf("no value with such ID")
	}

	return Data, nil
}

func (cs *InMemoryStorage) buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
