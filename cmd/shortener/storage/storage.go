package storage

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"github.com/PaBah/url-shortener.git/cmd/shortener/config"
	"hash/fnv"
	"os"
)

type Repository interface {
	Store(string) string
	FindByID(string) (string, error)
}

type InFileStorage struct {
	state map[string]string
}

func (cs *InFileStorage) Store(Data string) (ID string) {
	cs.loadState()
	ID = cs.buildID(Data)
	cs.state[ID] = Data
	cs.saveState()
	return
}

func (cs *InFileStorage) FindByID(ID string) (Data string, err error) {
	cs.loadState()
	Data = cs.state[ID]
	return Data, nil
}

func (cs *InFileStorage) loadState() {
	dat, err := os.ReadFile(config.StoragePath)
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

func (cs *InFileStorage) saveState() {
	b := &bytes.Buffer{}

	e := gob.NewEncoder(b)
	e.Encode(cs.state)

	os.WriteFile(config.StoragePath, b.Bytes(), 0644)
}

func (cs *InFileStorage) buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
