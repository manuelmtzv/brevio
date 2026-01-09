package ttl

import "time"

type Policy interface {
	TTL() time.Duration
}
