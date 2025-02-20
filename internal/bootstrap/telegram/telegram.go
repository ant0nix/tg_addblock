package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	Bot        *gotgbot.Bot
	Dispatcher *ext.Dispatcher
	Updater    *ext.Updater
}

func SetupBot(cfg *config.Config) (*Bot, error) {
	bot, err := gotgbot.NewBot(cfg.TgToken, nil)
	if err != nil {
		return nil, err
	}
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			logrus.Errorf("an error occurred while handling update: %v", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	return &Bot{Bot: bot, Dispatcher: dispatcher, Updater: updater}, nil
}

func (bot *Bot) Start() error {
	return bot.Updater.StartPolling(bot.Bot, nil)
}
