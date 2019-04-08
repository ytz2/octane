package handler

import (
	"golang.org/x/time/rate"
	"log"
	"octane/record"
)

type RateLimiter interface {
	Allow(k string) bool
}

type rateLimiter struct {
	limiters map[string]*rate.Limiter
}

// NewLotto ...
func NewRateLimiter(c *record.Config) RateLimiter {
	limiters := make(map[string]*rate.Limiter)
	for _, limit := range c.RateLimit {
		limiters[limit.Name] = rate.NewLimiter(rate.Limit(limit.Limit), limit.Burst)
		log.Printf("Creating ratelimit on %v with rate %v and burst %v", limit.Name, limit.Limit, limit.Burst)
	}
	return &rateLimiter{
		limiters: limiters,
	}
}

func (r *rateLimiter) Allow(k string) bool {
	l, ok := r.limiters[k]
	if !ok {
		return false
	}
	return l.Allow()
}
