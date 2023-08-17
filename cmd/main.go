package main

import (
	"context"
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/plinkplenk/simple_shortner/internal/api/handlers"
	"github.com/plinkplenk/simple_shortner/internal/api/routers"
	"github.com/plinkplenk/simple_shortner/internal/config"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var cfg = config.AppConfig

func init() {
	sql, err := goose.OpenDBWithDriver("pgx", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(sql, "migrations"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	pool, err := pgxpool.New(context.Background(), cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{ErrorHandler: handlers.APIErrorHandler})
	routers.Setup(app, pool)
	log.Fatal(app.Listen(cfg.Port))
}
