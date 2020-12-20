package cache_test

import (
	"testing"

	"github.com/onurcevik/rest-go/internal/model"

	"github.com/onurcevik/rest-go/internal/cache"
)

func TestCache(t *testing.T) {

	note := model.Note{
		ID:      900,
		Content: "testcontent",
	}
	cache := cache.NewRedisCache("localhost:6379", 0, 20)
	err := cache.Set("900", note)
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = cache.Get("900")
	if err != nil {
		t.Errorf(err.Error())
	}

}
