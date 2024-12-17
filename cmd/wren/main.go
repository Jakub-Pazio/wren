package main

import (
	"log/slog"

	"github.com/Jakub-Pazio/Wren/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("server failed", "error", err)
	}
}
