package main

import (
	"github.com/marmota-alpina/rate-limiter/internal/handler"
	"github.com/marmota-alpina/rate-limiter/internal/limiter"
	"github.com/marmota-alpina/rate-limiter/internal/middleware"
	"log"
	"net/http"

	"github.com/marmota-alpina/rate-limiter/config"
)

func main() {
	cfg := config.Load()
	store := limiter.NewRedisStorage(cfg.RedisHost, cfg.RedisPassword, cfg.RedisDB)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HealthHandler)

	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", middleware.RateLimiterMiddleware(cfg, store)(mux))
	if err != nil {
		log.Fatal(err)
	}
}
