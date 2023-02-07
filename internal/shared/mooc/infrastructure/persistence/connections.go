package persistence

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// NewRedisClient method that will connect a redis client and returns an instance of redis.Client
func NewRedisClient(config Configuration) (*redis.Client, error) {
	switch config.Type {
	case NoSQL:
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprint(config.Host, ":", config.Port),
			Password: config.Password,
			DB:       0,
		})

		return client, client.Ping(context.TODO()).Err()
	default:
		panic(fmt.Sprintf("%T type is not supported", config.Type))
	}
}

// NewRedisDataBase method that returns an instance of redis.Client
// if an error occurs a panic will be launched
func NewRedisDataBase(config Configuration) (db *redis.Client) {
	db, err := NewRedisClient(config)
	if err != nil {
		panic(err)
	}
	return
}
