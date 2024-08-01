package api

import (
	"context"
	"net/http"

	"preference-game/internal/service"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"handler",
		fx.Provide(
			NewHandler,
		),
		fx.Invoke(
			StartServer,
		),
	)
}

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) InitCard(_ http.ResponseWriter, _ *http.Request) {

}

func (h *Handler) OpenCard(_ http.ResponseWriter, _ *http.Request) {

}

func StartServer(lc fx.Lifecycle, h *Handler) {
	router := http.NewServeMux()

	router.HandleFunc("/initCard", h.InitCard)
	router.HandleFunc("/openCard", h.OpenCard)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
