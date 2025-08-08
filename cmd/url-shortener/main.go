package main

import (
	"fmt"

	"github.com/Danil-Zlo/url_shortener/internal/config"
)

func main() {
	// init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}

func setupLogger(env string) {
	//
}
