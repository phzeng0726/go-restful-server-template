package http

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/phzeng0726/go-server-template/internal/config"
)

func corsMiddleware() cors.Config {
	allowedOrigins := strings.Split(config.Env.AccessAllowOrigin, ",")
	for i, origin := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(origin)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowedOrigins
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Authorization", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Cache-Control", "X-Requested-With"}
	corsConfig.AllowCredentials = true

	return corsConfig
}
