package ratelimiter

import (
	"github.com/goal-web/contracts"
	"go.uber.org/ratelimit"
)

func Middleware(rate int, opts ...ratelimit.Option) interface{} {
	return func(request contracts.HttpRequest, pipe contracts.Pipe, limiter contracts.RateLimiter) interface{} {
		limiter.Limiter("request", func() contracts.Limiter {
			return ratelimit.New(rate, opts...)
		}).Take()

		return pipe(request)
	}
}
