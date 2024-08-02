package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type ReqCard struct {
	Suit      string `json:"suit"`
	Value     string `json:"value"`
	PromoCode string `json:"promo_code,omitempty"`
}

type Attempt struct {
	ID         uuid.UUID `json:"id"`
	UserID     string    `json:"user_id"`
	FirstCard  Card      `json:"first_card"`
	SecondCard Card      `json:"second_card"`
	IsWin      *bool     `json:"is_win"`
	PromoCode  string    `json:"promo_code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserIDCtxKey struct{}
