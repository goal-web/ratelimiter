package ratelimiter

import (
	"go.uber.org/ratelimit"
	"time"
)

type Limiter struct {
	limiter ratelimit.Limiter
}

func (l Limiter) Take() time.Time {
	return l.limiter.Take()
}
