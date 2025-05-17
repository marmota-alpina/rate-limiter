package limiter

import "time"

type Storage interface {
	Increment(key string, window time.Duration, maxRequests int) (int, error)
	Block(key string, duration time.Duration) error
	IsBlocked(key string) (bool, error)
}
