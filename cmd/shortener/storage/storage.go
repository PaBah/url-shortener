package storage

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"hash/fnv"
	"os"
)

type Repository interface {
	Store(string) string
	FindByID(string) (string, error)
}

type CringeStorage struct {
	state map[string]string
}

func (cs *CringeStorage) Store(Data string) (ID string) {
	cs.loadState()
	ID = cs.buildID(Data)
	cs.state[ID] = Data
	cs.saveState()
	return
}

func (cs *CringeStorage) FindByID(ID string) (Data string, err error) {
	cs.loadState()
	Data = cs.state[ID]
	return Data, nil
}

func (cs *CringeStorage) loadState() {
	dat, err := os.ReadFile("./.store")
	if err != nil {
		cs.state = make(map[string]string)
		return
	}

	var decodedState map[string]string
	d := gob.NewDecoder(bytes.NewReader(dat))

	err = d.Decode(&decodedState)
	if err != nil {
		cs.state = make(map[string]string)
	}
	cs.state = decodedState
}

func (cs *CringeStorage) saveState() {
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)
	e.Encode(cs.state)

	os.WriteFile("./.store", b.Bytes(), 0644)
}

func (cs *CringeStorage) buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
