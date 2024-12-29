package main

import (
	"log/slog"

	"github.com/Jakub-Pazio/wren/internal/server"
)

func main() {
	srv := server.New()
	if err := srv.Run(); err != nil {
		slog.Error("server failed", "error", err)
	}
}
