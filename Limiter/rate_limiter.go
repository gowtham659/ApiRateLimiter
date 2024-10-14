package limiter

import (
    "sync"
    "time"
)

// TokenBucket represents a rate limiter using the token bucket algorithm
type TokenBucket struct {
    Tokens     int           // Current number of tokens
    MaxTokens  int           // Maximum capacity of tokens
    RefillRate time.Duration // Rate at which tokens refill
    LastRefill time.Time     // Last time the bucket was refilled
    mu         sync.Mutex    // Mutex to protect concurrent access
}

// NewTokenBucket creates a new TokenBucket
func NewTokenBucket(maxTokens int, refillRate time.Duration) *TokenBucket {
    return &TokenBucket{
        Tokens:     maxTokens,
        MaxTokens:  maxTokens,
        RefillRate: refillRate,
        LastRefill: time.Now(),
    }
}

// Allow checks if a request can proceed by consuming a token
func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()

    // Refill tokens based on time elapsed
    now := time.Now()
    timeSinceLastRefill := now.Sub(tb.LastRefill)
    tokensToAdd := int(timeSinceLastRefill / tb.RefillRate)

    if tokensToAdd > 0 {
        tb.Tokens = min(tb.Tokens+tokensToAdd, tb.MaxTokens)
        tb.LastRefill = now
    }

    // If tokens are available, allow the request and consume a token
    if tb.Tokens > 0 {
        tb.Tokens--
        return true
    }

    return false
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
