package storage

import (
	"encoding/hex"
	"hash/fnv"
)

type Repository interface {
	Store(string) string
	FindByID(string) (string, error)
}

type InFileStorage struct {
	state map[string]string
}

func (cs *InFileStorage) Store(Data string) (ID string) {
	if cs.state == nil {
		cs.state = make(map[string]string)
	}

	ID = cs.buildID(Data)
	cs.state[ID] = Data
	return
}

func (cs *InFileStorage) FindByID(ID string) (Data string, err error) {
	if cs.state == nil {
		cs.state = make(map[string]string)
	}

	Data = cs.state[ID]
	return Data, nil
}

func (cs *InFileStorage) buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
