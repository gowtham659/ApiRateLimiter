package middleware

import (
    "net/http"
    "sync"
    "time"
    "ApiRateLimiterApp/Limiter"
)

var (
    buckets = make(map[string]*limiter.TokenBucket) // Map of user buckets
    mu      sync.Mutex
)

// RateLimitMiddleware applies rate limiting to incoming requests
func RateLimitMiddleware(maxTokens int, refillRate time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            userID := r.RemoteAddr // Simple way to identify a user, can be enhanced (e.g., API keys)

            mu.Lock()
            bucket, exists := buckets[userID]
            if !exists {
                bucket = limiter.NewTokenBucket(maxTokens, refillRate)
                buckets[userID] = bucket
            }
            mu.Unlock()

            if bucket.Allow() {
                next.ServeHTTP(w, r) // Allow the request if bucket has tokens
            } else {
                http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
            }
        })
    }
}
