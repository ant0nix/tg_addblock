package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func InitCommands(dsp *ext.Dispatcher) {
	dsp.AddHandler(handlers.NewCommand("start", start))
	return
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(ctx.EffectiveMessage.Chat.Id, fmt.Sprintf("Hello!"), nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
