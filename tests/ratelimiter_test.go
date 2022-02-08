package tests

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/ratelimiter"
	"go.uber.org/ratelimit"
	"log"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	rate := ratelimiter.RateLimiter{}
	prev := time.Now()

	for i := 0; i < 20; i++ {
		now := rate.Limiter("testing", func() contracts.Limiter {
			return ratelimit.New(10, ratelimit.Per(time.Second)) // per second
		}).Take()
		log.Default().Println(i, now.Sub(prev))
		prev = now
	}
}

func BenchmarkLimiter(b *testing.B) {
	rate := ratelimiter.RateLimiter{}
	per := b.N / 2
	if per <= 0 {
		per = 1
	}

	for i := 0; i < b.N; i++ {
		rate.Limiter("testing", func() contracts.Limiter {
			return ratelimit.New(per, ratelimit.Per(time.Second)) // per second
		}).Take()
	}
}
