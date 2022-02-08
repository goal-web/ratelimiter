package ratelimiter

import (
	"github.com/goal-web/contracts"
	"sync"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application contracts.Application) {
	application.Singleton("ratelimiter", func() contracts.RateLimiter {
		return &RateLimiter{limiters: sync.Map{}}
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
