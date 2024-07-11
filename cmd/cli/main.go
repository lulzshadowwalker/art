package main

import (
	"log/slog"
	"os"

	"github.com/lulzshadowwalker/art/internal/cli"
)

func main() {
	slog.Info("cli starting", "args", os.Args[1:])
	if err := cli.Run(os.Args[1:]); err != nil {
		slog.Error("cli failure", "err", err)
		os.Exit(1)
	}

	slog.Info("cli exiting")
}
