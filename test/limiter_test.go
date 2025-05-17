package test

import (
	"testing"
	"time"
)

type MemoryStore struct {
	data map[string]int
}

func (m *MemoryStore) Increment(key string, window time.Duration, max int) (int, error) {
	m.data[key]++
	return m.data[key], nil
}

func (m *MemoryStore) Block(key string, duration time.Duration) error {
	return nil
}

func (m *MemoryStore) IsBlocked(key string) (bool, error) {
	return false, nil
}

func TestLimiterMemory(t *testing.T) {
	mem := &MemoryStore{data: map[string]int{}}

	count, _ := mem.Increment("ip:127.0.0.1", time.Minute, 5)
	if count != 1 {
		t.Fatal("expected 1 request count")
	}
}
