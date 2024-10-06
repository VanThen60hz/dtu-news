package database

import (
	"context"
	"fmt"
	"log"

	"gonews/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDB(config *config.Config) *redis.Client {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Connected successfully to Redis!")
	}

	return rdb
}
