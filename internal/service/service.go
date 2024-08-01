package service

import (
	"math/rand"
	"time"

	"preference-game/internal/entity"
	"preference-game/internal/repository"

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
var AllSuitMap = map[string]struct{}{
	"spades":   {},
	"clubs":    {},
	"diamonds": {},
	"hearts":   {},
}
var AllValuesCard = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var AllValuesCardMap = map[string]struct{}{
	"2":  {},
	"3":  {},
	"4":  {},
	"5":  {},
	"6":  {},
	"7":  {},
	"8":  {},
	"9":  {},
	"10": {},
	"J":  {},
	"Q":  {},
	"K":  {},
	"A":  {},
}

func (s *Service) InitCard() entity.Card {
	rand.NewSource(time.Now().UnixNano())
	suit := AllSuit[rand.Intn(len(AllSuit))]
	value := AllValuesCard[rand.Intn(len(AllValuesCard)-1)+1] // Avoid "A"
	return entity.Card{Suit: suit, Value: value}
}

// todo: 33% win
func (s *Service) getRandomCard(firstCard entity.Card) entity.Card {
	return entity.Card{}
}

func (s *Service) OpenCard(firstCard entity.Card) (entity.Card, bool) {
	secondCard := s.getRandomCard(firstCard)
	isWin := checkWin(firstCard, secondCard)
	return secondCard, isWin
}

func checkWin(firstCard, secondCard entity.Card) bool {
	firstValueIndex := getValueIndex(firstCard.Value)
	secondValueIndex := getValueIndex(secondCard.Value)
	return secondValueIndex > firstValueIndex
}

func getValueIndex(value string) int {
	for i, v := range AllValuesCard {
		if v == value {
			return i
		}
	}
	return -1
}
