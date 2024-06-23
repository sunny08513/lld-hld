package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity       int
	tokens         int
	refillRate     int
	refillInterval time.Duration
	ticker         *time.Ticker
	mutex          *sync.Mutex
}

func NewTokenBucket(capacity, refillRate int, refillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		refillRate:     refillRate,
		refillInterval: refillInterval,
		mutex:          &sync.Mutex{},
	}

	tb.startRefilling()
	return tb
}

func (tb *TokenBucket) startRefilling() {
	tb.ticker = time.NewTicker(tb.refillInterval)
	go func() {
		for range tb.ticker.C {
			tb.mutex.Lock()
			tb.tokens += tb.refillRate
			if tb.tokens > tb.capacity {
				tb.tokens = tb.capacity
			}

			tb.mutex.Unlock()
		}
	}()
}

func (tb *TokenBucket) ALlow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func main() {
	tb := NewTokenBucket(10, 1, time.Second)
	for i := 1; i <= 20; i++ {
		if tb.ALlow() {
			fmt.Printf("%v Request Allowed\n", i)
		} else {
			fmt.Printf("%v Request Not Allowed\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
