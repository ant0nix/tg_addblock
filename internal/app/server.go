package app

import (
	"context"
	"github.com/ant0nix/vpn.git/internal/app/handler"
	"github.com/ant0nix/vpn.git/internal/bootstrap"
	"github.com/ant0nix/vpn.git/internal/config"
	"github.com/ant0nix/vpn.git/internal/repositories/cache"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bot, err := bootstrap.SetupBot(cfg)
	if err != nil {
		return err
	}

	cash := cacheclient.Init()
	h := handler.InitHandler(cash)
	h.InitCommands(cfg, bot.Dispatcher)
	err = bot.Start()
	if err != nil {
		return err
	}
	logrus.Info("Starting poling")
	bot.Updater.Idle()
	return nil
}
