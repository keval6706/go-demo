package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Redis Connected.")

	return client
}
