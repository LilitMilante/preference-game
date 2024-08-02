package repository

import (
	"context"
	"time"

	"preference-game/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
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
	c  *cache.Cache
}

func NewRepository(db *pgxpool.Pool) *Repository {
	expiration := time.Until(time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour))
	c := cache.New(expiration, 24*time.Hour)
	return &Repository{
		db: db,
		c:  c,
	}
}

func (r *Repository) SaveAttempt(ctx context.Context, a entity.Attempt) error {
	q := `
		INSERT INTO attempts (id, user_id, first_card_suit, first_card_value, second_card_suit, second_card_value, is_win, promo_code, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        `

	_, err := r.db.Exec(
		ctx,
		q,
		a.ID,
		a.UserID,
		a.FirstCard.Suit,
		a.FirstCard.Value,
		a.SecondCard.Suit,
		a.SecondCard.Value,
		a.IsWin,
		a.PromoCode,
		a.CreatedAt,
		a.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatedAttempt(ctx context.Context, a entity.Attempt) error {
	q := `
		UPDATE attempts
		SET second_card_suit = $1, second_card_value = $2, is_win = $3, promo_code = $4, updated_at = $5
		WHERE id = $6
`

	_, err := r.db.Exec(ctx, q, a.SecondCard.Suit, a.SecondCard.Value, a.IsWin, a.PromoCode, a.UpdatedAt, a.ID)
	if err != nil {
		return err
	}

	r.c.Set(a.ID.String(), a, 0)

	return nil
}

func (r *Repository) Attempt(ctx context.Context, userID string) (a entity.Attempt, err error) {
	q := `
		SELECT id, user_id, first_card_suit, first_card_value, second_card_suit, second_card_value, is_win, promo_code, created_at, updated_at
		FROM attempts
		WHERE user_id = $1 AND is_win IS NULL
`
	err = r.db.QueryRow(ctx, q, userID).
		Scan(
			&a.ID,
			&a.UserID,
			&a.FirstCard.Suit,
			&a.FirstCard.Value,
			&a.SecondCard.Suit,
			&a.SecondCard.Value,
			&a.IsWin,
			&a.PromoCode,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
	if err != nil {
		return entity.Attempt{}, err
	}

	return a, nil
}
