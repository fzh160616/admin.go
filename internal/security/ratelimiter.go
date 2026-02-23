package security

import (
	"fmt"
	"sync"
	"time"
)

type entry struct {
	Count      int
	FirstAt    time.Time
	BlockedTil time.Time
}

type LoginRateLimiter struct {
	mu             sync.Mutex
	store          map[string]*entry
	maxAttempts    int
	window         time.Duration
	blockDuration  time.Duration
}

func NewLoginRateLimiter(maxAttempts int, window, blockDuration time.Duration) *LoginRateLimiter {
	return &LoginRateLimiter{
		store:         make(map[string]*entry),
		maxAttempts:   maxAttempts,
		window:        window,
		blockDuration: blockDuration,
	}
}

func (r *LoginRateLimiter) key(ip, account string) string {
	return fmt.Sprintf("%s|%s", ip, account)
}

func (r *LoginRateLimiter) Allow(ip, account string) (bool, time.Duration) {
	now := time.Now()
	k := r.key(ip, account)

	r.mu.Lock()
	defer r.mu.Unlock()

	e, ok := r.store[k]
	if !ok {
		return true, 0
	}

	if now.Before(e.BlockedTil) {
		return false, time.Until(e.BlockedTil)
	}

	if now.Sub(e.FirstAt) > r.window {
		delete(r.store, k)
		return true, 0
	}

	return true, 0
}

func (r *LoginRateLimiter) RecordFailure(ip, account string) {
	now := time.Now()
	k := r.key(ip, account)

	r.mu.Lock()
	defer r.mu.Unlock()

	e, ok := r.store[k]
	if !ok || now.Sub(e.FirstAt) > r.window {
		r.store[k] = &entry{Count: 1, FirstAt: now}
		return
	}

	e.Count++
	if e.Count >= r.maxAttempts {
		e.BlockedTil = now.Add(r.blockDuration)
	}
}

func (r *LoginRateLimiter) RecordSuccess(ip, account string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.store, r.key(ip, account))
}
