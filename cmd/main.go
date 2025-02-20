package main

import (
	"context"
	"github.com/ant0nix/tg_addblock/internal/app"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("failed to read config:", err)
	}

	logrus.Infof("Debug level: %v", cfg.LogLevel)

	err = app.Run(ctx, cfg)
	if err != nil {
		logrus.Fatal("Critical error:", err)
	}
}
