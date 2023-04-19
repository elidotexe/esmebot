package storage

import "sync"

type Storage interface {
	Add(c int64, u int64, i Info)
	// Remove()
	// Exist()
}

func NewMemoryStorage() *MemoryStorage {
	m := MemoryStorage{}
	m.pending = make(map[string]Info)
	return &m
}

type MemoryStorage struct {
	sync.RWMutex
	pending map[string]Info
}

type Info struct {
	VerifyMsg string
	IsHuman   bool
}

func (m *MemoryStorage) Add(c int64, u int64, i Info) {
	m.Lock()
	defer m.Unlock()
}
