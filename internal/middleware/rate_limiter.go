package middleware

import (
	"github.com/marmota-alpina/rate-limiter/config"
	"github.com/marmota-alpina/rate-limiter/internal/limiter"
	"net"
	"net/http"
	"time"
)

func RateLimiterMiddleware(cfg *config.Config, store limiter.Storage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			key := ip

			if key == net.IPv6loopback.String() {
				key = "127.0.0.1"
			}

			limit := cfg.PerIP
			window := cfg.BlockDurationIP

			if token := r.Header.Get("API_KEY"); token != "" {
				key = "token:" + token
				limit = cfg.PerToken
				window = cfg.BlockDurationToken
			}

			if blocked, _ := store.IsBlocked(key); blocked {
				http.Error(w, `{"error":"you have reached the maximum number of requests or actions allowed within a certain time frame"}`, http.StatusTooManyRequests)
				return
			}

			count, _ := store.Increment(key, time.Second, limit)
			if count > limit {
				err := store.Block(key, window)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Error(w, `{"error":"you have reached the maximum number of requests or actions allowed within a certain time frame"}`, http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
