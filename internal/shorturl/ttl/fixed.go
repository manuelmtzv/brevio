package ttl

import "time"

type Fixed struct {
	value time.Duration
}

func NewFixed(d time.Duration) Fixed {
	return Fixed{value: d}
}

func (f Fixed) TTL() time.Duration {
	return f.value
}
