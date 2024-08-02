package service

import (
	"context"
	"crypto/rand"
	"math/big"
	"slices"
	"time"

	"preference-game/internal/entity"
	"preference-game/internal/repository"

	"github.com/gofrs/uuid"
	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(
			NewService,
		),
	)
}

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

var AllSuit = []string{"spades", "clubs", "diamonds", "hearts"}

var AllValuesCard = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func (s *Service) InitCard(ctx context.Context) (entity.Card, error) {
	suitIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllSuit))))
	if err != nil {
		return entity.Card{}, err
	}
	suit := AllSuit[suitIndex.Int64()]

	valueIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllValuesCard)-1)))
	if err != nil {
		return entity.Card{}, err
	}
	value := AllValuesCard[(valueIndex.Int64() + 1)]

	firstCard := entity.Card{Suit: suit, Value: value}

	attempt := entity.Attempt{
		ID:        uuid.Must(uuid.NewV4()),
		UserID:    ctx.Value(entity.UserIDCtxKey{}).(string),
		FirstCard: firstCard,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.SaveAttempt(ctx, attempt)
	if err != nil {
		return entity.Card{}, err
	}

	return firstCard, nil
}

func (s *Service) OpenCard(ctx context.Context) (entity.ReqCard, error) {
	userID := ctx.Value(entity.UserIDCtxKey{}).(string)

	attempt, err := s.repo.Attempt(ctx, userID)
	if err != nil {
		return entity.ReqCard{}, err
	}

	isWin, err := calculateWin()
	if err != nil {
		return entity.ReqCard{}, err
	}

	if isWin {
		attempt.PromoCode = "CARD50"
		attempt.SecondCard, err = s.randWinCard(attempt.FirstCard)
	} else {
		attempt.SecondCard, err = s.randLoseCard(attempt.FirstCard)
	}
	if err != nil {
		return entity.ReqCard{}, err
	}

	attempt.IsWin = &isWin
	attempt.UpdatedAt = time.Now()

	err = s.repo.UpdatedAttempt(ctx, attempt)
	if err != nil {
		return entity.ReqCard{}, err
	}

	req := entity.ReqCard{
		Suit:      attempt.SecondCard.Suit,
		Value:     attempt.SecondCard.Value,
		PromoCode: attempt.PromoCode,
	}

	return req, nil
}

func calculateWin() (bool, error) {
	const winValue = 1

	randValue, err := rand.Int(rand.Reader, big.NewInt(3))
	if err != nil {
		return false, err
	}

	return randValue.Int64() == winValue, nil
}

func (s *Service) randWinCard(initCard entity.Card) (entity.Card, error) {
	suitIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllSuit))))
	if err != nil {
		return entity.Card{}, err
	}
	suit := AllSuit[suitIndex.Int64()]

	initCardIndex := slices.Index(AllValuesCard, initCard.Value)

	valueIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllValuesCard)-initCardIndex)))
	if err != nil {
		return entity.Card{}, err
	}
	value := AllValuesCard[(valueIndex.Int64() + int64(initCardIndex))]

	card := entity.Card{Suit: suit, Value: value}

	return card, nil
}

func (s *Service) randLoseCard(initCard entity.Card) (entity.Card, error) {
	suitIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllSuit))))
	if err != nil {
		return entity.Card{}, err
	}
	suit := AllSuit[suitIndex.Int64()]

	initCardIndex := slices.Index(AllValuesCard, initCard.Value)

	valueIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(AllValuesCard[:initCardIndex]))))
	if err != nil {
		return entity.Card{}, err
	}
	value := AllValuesCard[(valueIndex.Int64())]

	card := entity.Card{Suit: suit, Value: value}

	return card, nil
}
