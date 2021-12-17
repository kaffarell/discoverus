package adapters

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func CreateClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
