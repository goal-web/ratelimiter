package ratelimiter

import (
	"github.com/goal-web/contracts"
	"sync"
)

type RateLimiter struct {
	limiters sync.Map
}

func (rate *RateLimiter) Limiter(name string, limiter func() contracts.Limiter) contracts.Limiter {
	if limit, exists := rate.limiters.Load(name); exists {
		return limit.(contracts.Limiter)
	}
	limit := limiter()
	rate.limiters.Store(name, limit)
	return limit
}
