package repository

import (
	"context"
	"time"

	"preference-game/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"repository",
		fx.Provide(
			NewRepository,
		),
	)
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveAttempt(ctx context.Context, attempt entity.Attempt) error {
	q := `INSERT INTO attempts (user_id, first_card_suit, first_card_value, second_card_suit, second_card_value, is_win, promocode, attempt_date) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(ctx, q,
		attempt.UserID, attempt.FirstCard.Suit, attempt.FirstCard.Value, attempt.SecondCard.Suit, attempt.SecondCard.Value, attempt.IsWin, attempt.Promocode, time.Now(),
	)
	return err
}
