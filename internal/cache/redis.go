package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/onurcevik/rest-go/internal/model"
)

type NotesCache interface {
	Set(key string, value model.Note) error
	Get(key string) (*model.Note, error)
}

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) NotesCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

//NewRedisClient dials Redis cache and returns client
func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "", // no password set
		DB:       cache.db,
	})

}

//Ping can be used to test connection
func Ping(client *redis.Client) (string, error) {
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return "", err
	}
	return pong, nil

}

func (cache *redisCache) Set(key string, value model.Note) error {

	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	client.Set(client.Context(), key, json, cache.expires*time.Minute)
	return nil
}
func (cache *redisCache) Get(key string) (*model.Note, error) {

	client := cache.getClient()
	val, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return nil, err
	}
	note := model.Note{}
	err = json.Unmarshal([]byte(val), &note)
	if err != nil {
		return nil, err
	}
	return &note, nil

}
