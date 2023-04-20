package storage

import (
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Storage interface {
	Add(c int64, u int64, i Info)
	Remove(c int64, u int64)
	Exist(c int64, u int64) (Info, bool)
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
	CaptchaMessage tgbotapi.Message
	IsHuman        bool
}

func (m *MemoryStorage) Add(c int64, u int64, i Info) {
	m.Lock()
	defer m.Unlock()
	m.pending[key(c, u)] = i
}

func (m *MemoryStorage) Remove(c int64, u int64) {
	m.Lock()
	defer m.Unlock()
	delete(m.pending, key(c, u))
}

func (m *MemoryStorage) Exist(c int64, u int64) (Info, bool) {
	m.RLock()
	defer m.RUnlock()
	i, ok := m.pending[key(c, u)]
	return i, ok
}

func key(c int64, u int64) string {
	return fmt.Sprintf("%d-%d", c, u)
}
