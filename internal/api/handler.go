package api

import (
	"context"
	"encoding/json"
	"net/http"

	"preference-game/internal/entity"
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

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.Header.Get("X-UserId")
		if id == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, entity.UserIDCtxKey{}, id)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func (h *Handler) InitCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	card, err := h.s.InitCard(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type Request struct {
	Card      entity.Card
	PromoCode string `json:"promo_code"`
}

func (h *Handler) OpenCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	card, err := h.s.OpenCard(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func StartServer(lc fx.Lifecycle, h *Handler) {
	router := http.NewServeMux()

	router.HandleFunc("/initCard", AuthMiddleware(h.InitCard))
	router.HandleFunc("/openCard", AuthMiddleware(h.OpenCard))

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
