package store

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/manuelmtzv/brevio/internal/models"
)

type ShortURLStore struct {
	rc *redis.Client
}

func NewShortURLStore(rc *redis.Client) *ShortURLStore {
	return &ShortURLStore{
		rc: rc,
	}
}

func (s *ShortURLStore) Create(ctx context.Context, data models.CreateShortURL) (*models.ShortURL, error) {
	key := s.key(data.Code)

	exists, err := s.rc.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, ErrCodeAlreadyExists
	}

	createdAt := time.Now()

	fields := map[string]any{
		"code":      data.Code,
		"target":    data.Target,
		"visits":    0,
		"createdAt": createdAt.Unix(),
	}

	if err := s.rc.HSet(ctx, key, fields).Err(); err != nil {
		return nil, err
	}

	if data.TTL != nil {
		if err := s.rc.Expire(ctx, key, *data.TTL).Err(); err != nil {
			return nil, err
		}
	}

	return &models.ShortURL{
		Code:      data.Code,
		Target:    data.Target,
		Visits:    0,
		CreatedAt: createdAt,
	}, nil
}

func (s *ShortURLStore) FindByCode(ctx context.Context, code string) (*models.ShortURL, error) {
	key := s.key(code)

	values, err := s.rc.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}

	visits, err := strconv.Atoi(values["visits"])
	if err != nil {
		return nil, err
	}

	createdAtUnix, err := strconv.ParseInt(values["createdAt"], 10, 64)
	if err != nil {
		return nil, err
	}

	return &models.ShortURL{
		Code:      values["code"],
		Target:    values["target"],
		Visits:    visits,
		CreatedAt: time.Unix(createdAtUnix, 0),
	}, nil
}

func (s *ShortURLStore) IncrementVisits(ctx context.Context, code string) error {
	key := s.key(code)
	return s.rc.HIncrBy(ctx, key, "visits", 1).Err()
}

func (s *ShortURLStore) key(code string) string {
	return "shorturl:" + code
}
