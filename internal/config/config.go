package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type config struct {
	DbURL     string `env:"DB_URL" env-default:"postgresql://postgres:postgres@localhost:5432/postgres"`
	Port      string `env:"PORT" env-default:":3000"`
	JWTSecret string `env:"SECRET" env-default:"secret"`
}

var AppConfig config

func init() {
	_ = godotenv.Load()
	err := cleanenv.ReadEnv(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
