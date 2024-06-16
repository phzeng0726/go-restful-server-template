package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/phzeng0726/go-server-template/internal/config"
)

func main() {
	router.Run(config.Env.Host + ":" + config.Env.Port)
}
