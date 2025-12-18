package store

import "github.com/go-redis/redis/v8"

type Storage struct {
	ShortUrls ShortURLStorage
}

func NewStorage(rc *redis.Client) *Storage {
	return &Storage{
		ShortUrls: NewShortURLStore(rc),
	}
}
