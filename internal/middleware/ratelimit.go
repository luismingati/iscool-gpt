package middleware

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	maxReqs  int
	window   time.Duration
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		maxReqs:  maxRequests,
		window:   window,
	}

	// Start background cleanup goroutine
	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get client IP
		ip := r.RemoteAddr

		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()

		// Get existing requests for this IP
		reqs, exists := rl.requests[ip]
		if !exists {
			reqs = []time.Time{}
		}

		// Filter out requests outside the time window
		validReqs := []time.Time{}
		for _, reqTime := range reqs {
			if now.Sub(reqTime) < rl.window {
				validReqs = append(validReqs, reqTime)
			}
		}

		// Check if limit exceeded
		if len(validReqs) >= rl.maxReqs {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error: "Rate limit exceeded. Try again later.",
			})
			return
		}

		// Add current request
		validReqs = append(validReqs, now)
		rl.requests[ip] = validReqs

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

// cleanup removes old entries periodically
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()

		for ip, reqs := range rl.requests {
			validReqs := []time.Time{}
			for _, reqTime := range reqs {
				if now.Sub(reqTime) < rl.window {
					validReqs = append(validReqs, reqTime)
				}
			}

			if len(validReqs) == 0 {
				delete(rl.requests, ip)
			} else {
				rl.requests[ip] = validReqs
			}
		}
		rl.mu.Unlock()
	}
}
