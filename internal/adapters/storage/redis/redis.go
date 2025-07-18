package redis

import (
	"context"
	"fmt"

	"github.com/mrfade/case-sss/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type DatabaseType int

const (
	DBCache     DatabaseType = iota // 1
	DBRateLimit                     // 2
)

type StorageClient struct {
	connections map[DatabaseType]*redis.Client
}

type Storage struct {
	Client *redis.Client
}

func NewClient(host, port, password string) *StorageClient {
	addr := fmt.Sprintf("%s:%s", host, port)

	var connections = make(map[DatabaseType]*redis.Client)
	for db := DBCache; db <= DBRateLimit; db++ {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       int(db),
		})

		_, err := client.Ping(context.TODO()).Result()
		if err != nil {
			panic(errors.ErrUnableToConnectRedis)
		}

		connections[db] = client
	}

	return &StorageClient{
		connections: connections,
	}
}

func (sc *StorageClient) ForDatabase(db DatabaseType) *Storage {
	if client, ok := sc.connections[db]; ok {
		return &Storage{
			Client: client,
		}
	}
	return nil
}

func (sc *StorageClient) Close() error {
	var lastErr error

	for dbType, client := range sc.connections {
		if err := client.Close(); err != nil {
			lastErr = err
			fmt.Printf("error closing redis connection for %d: %v\n", dbType, err)
		}
	}

	return lastErr
}
