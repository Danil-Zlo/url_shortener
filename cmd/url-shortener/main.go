package main

import (
	"fmt"
	"log"
	slog "log/slog"
	"os"

	"github.com/Danil-Zlo/url_shortener/internal/config"
	mwLogger "github.com/Danil-Zlo/url_shortener/internal/http-server/middleware/logger"
	sl "github.com/Danil-Zlo/url_shortener/internal/lib/logger/slog"
	"github.com/Danil-Zlo/url_shortener/internal/storage/sqlite"
	middleware "github.com/go-chi/chi/middleware"
	chi "github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// add enviroment var
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg.StoragePath)

	// init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	id, err := storage.SaveURL("https://ya.ru", "yandex")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}

	log.Info("save url", slog.Int64("id", id))

	_ = storage

	// init router: chi, "chi render"
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}
	return log
}
