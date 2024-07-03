package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	limit        int
	requests     []time.Time
	windowPeriod time.Duration
	mu           sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, windowPeriod time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:     []time.Time{},
		limit:        limit,
		windowPeriod: windowPeriod,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutOff := now.Add(-rl.windowPeriod)

	index := 0
	for i, request := range rl.requests {
		if request.After(cutOff) {
			index = i
			break
		}
	}

	rl.requests = rl.requests[index:]

	if len(rl.requests) < rl.limit {
		rl.requests = append(rl.requests, now)
		return true
	}

	return false
}

func main() {
	// Create a rate limiter that allows 5 requests per 10 seconds
	limiter := NewRateLimiter(5, 10*time.Second)

	// Simulate a series of requests
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request rate-limited")
		}
		time.Sleep(1 * time.Second)
	}
}
