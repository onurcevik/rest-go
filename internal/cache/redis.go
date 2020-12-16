package cache

import (
	"github.com/go-redis/redis/v8"
)

//Dials Redis cache and returns client
func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func Ping(client *redis.Client) (string, error) {
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return "", err
	}
	return pong, nil
	var d interface{}
}
