package config

import (
	"os"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	SecretKey      string
	ExpireDuration int64
}

var JwtConfigAuth *JwtConfig

func init() {
	godotenv.Load()

	JwtConfigAuth = &JwtConfig{
		SecretKey:      os.Getenv("JWT_SECRET_KEY"),
		ExpireDuration: 3600,
	}
}
