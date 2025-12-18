package models

import "time"

type ShortURL struct {
	Code      string    `json:"code"`
	Target    string    `json:"target"`
	Visits    int       `json:"visits"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateShortURL struct {
	Code   string         `json:"code"`
	Target string         `json:"target"`
	TTL    *time.Duration `json:"ttl"`
}
