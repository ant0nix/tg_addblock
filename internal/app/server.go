package app

import (
	"context"
	"github.com/ant0nix/tg_addblock/internal/app/handler"
	redisclient "github.com/ant0nix/tg_addblock/internal/bootstrap/redis"
	"github.com/ant0nix/tg_addblock/internal/bootstrap/telegram"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/ant0nix/tg_addblock/internal/repositories/cache"
	"github.com/sirupsen/logrus"
	"log"
)

func Run(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bot, err := telegram.SetupBot(cfg)
	if err != nil {
		return err
	}

	r := redisclient.New(cfg)
	go func() {
		err = r.Alive(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	h := handler.InitHandler(cacheclient.Init(), cacheclient.NewRedisRepository(r))
	h.InitCommands(cfg, bot.Dispatcher)
	err = bot.Start()
	if err != nil {
		return err
	}
	logrus.Info("Starting poling")
	bot.Updater.Idle()
	return nil
}
