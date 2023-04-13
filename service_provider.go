package ratelimiter

import (
	"github.com/goal-web/contracts"
	"sync"
)

type ServiceProvider struct {
}

func NewService() contracts.ServiceProvider {
	return &ServiceProvider{}
}

func (provider ServiceProvider) Register(application contracts.Application) {
	application.Singleton("ratelimiter", func() contracts.RateLimiter {
		return &RateLimiter{limiters: sync.Map{}}
	})
}

func (provider ServiceProvider) Start() error {
	return nil
}

func (provider ServiceProvider) Stop() {
}
