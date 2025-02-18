package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/ant0nix/vpn.git/internal/config"
	"github.com/ant0nix/vpn.git/internal/repositories"
)

type Handler struct {
	Cache repositories.CacheInterface
}

func InitHandler(cache repositories.CacheInterface) *Handler {
	return &Handler{Cache: cache}
}

func (h *Handler) InitCommands(cfg *config.Config, dsp *ext.Dispatcher) {

	a := func(msg *gotgbot.Message) bool {
		if msg.Chat.Id == cfg.ChatID && msg.MessageThreadId != 0 {
			return true
		}
		return false
	}

	dsp.AddHandler(handlers.NewCommand("start", h.start))
	dsp.AddHandler(handlers.NewMessage(a, h.handleNewComment))
}

func (h *Handler) start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(ctx.EffectiveMessage.Chat.Id, h.Cache.GetKeys(), nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
func (h *Handler) handleNewComment(b *gotgbot.Bot, ctx *ext.Context) error {
	h.Cache.Add(ctx.EffectiveMessage.MessageThreadId, ctx)
	return nil
}
