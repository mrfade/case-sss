package ports

import "time"

type Cacher interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (any, error)
	Del(key string) error
}
