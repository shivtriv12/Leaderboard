package redisClient

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

func Init() {
	once.Do(func() {
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			log.Fatal("redis related env not set")
		}
		client = redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		})
		if err := client.Ping(context.Background()).Err(); err != nil {
			log.Fatal(err)
		}
	})
}

func Get() *redis.Client {
	if client == nil {
		log.Fatal("redis client not initialized")
	}
	return client
}

func Close() error {
	if client != nil {
		log.Println("closing redis client")
		return client.Close()
	}
	return nil
}
