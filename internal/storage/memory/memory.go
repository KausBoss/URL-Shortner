package memory

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Store struct {
	data map[string]string
	sync.RWMutex
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (m *Store) Save(url string) (string, error) {
	m.Lock()
	defer m.Unlock()
	code := generateCode(url)
	m.data[code] = url
	return code, nil
}

func (m *Store) Load(code string) (string, error) {
	m.RLock()
	defer m.RUnlock()
	if url, exists := m.data[code]; exists {
		return url, nil
	}
	return "", errors.New("no URL Found")
}

func generateCode(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
}
