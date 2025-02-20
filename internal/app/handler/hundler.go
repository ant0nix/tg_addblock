package handler

import (
	"context"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/ant0nix/tg_addblock/internal/repositories"
)

type Handler struct {
	InMemCache repositories.CacheInterface
	Redis      repositories.RedisInterface
}

func InitHandler(cache repositories.CacheInterface, redis repositories.RedisInterface) *Handler {
	return &Handler{InMemCache: cache, Redis: redis}
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
	_, err := b.SendMessage(ctx.EffectiveMessage.Chat.Id, h.InMemCache.GetKeys(), nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func (h *Handler) handleNewComment(b *gotgbot.Bot, ctx *ext.Context) error {
	//todo recovery
	amount, err := h.Redis.Get(context.Background(), fmt.Sprintln(ctx.EffectiveMessage.MessageThreadId))
	if err != nil {
		return err
	}
	h.InMemCache.Add(ctx.EffectiveMessage.MessageThreadId, ctx)
	return h.Redis.Set(context.Background(), fmt.Sprintln(ctx.EffectiveMessage.MessageThreadId), amount+1)
}
