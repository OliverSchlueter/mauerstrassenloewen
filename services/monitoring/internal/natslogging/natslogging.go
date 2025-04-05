package natslogging

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log/slog"
)

type Handler struct {
	nats *nats.Conn
}

type Configuration struct {
	Nats *nats.Conn
}

func NewHandler(cfg Configuration) *Handler {
	return &Handler{
		nats: cfg.Nats,
	}
}

func (h *Handler) Register() error {
	if _, err := h.nats.Subscribe(">", h.handleMessage); err != nil {
		return fmt.Errorf("could not subscribe to NATS: %w", err)
	}

	return nil
}

func (h *Handler) handleMessage(msg *nats.Msg) {
	slog.Info(
		"Received NATS message",
		slog.String("subject", msg.Subject),
		slog.String("reply", msg.Reply),
		slog.String("data", string(msg.Data)),
	)
}
