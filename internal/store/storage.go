package store

import "github.com/go-redis/redis/v8"

type Storage struct {
	ShortURLs ShortURLStorage
}

func NewStorage(rc *redis.Client) Storage {
	return Storage{
		ShortURLs: NewShortURLStore(rc),
	}
}
