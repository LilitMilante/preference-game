package entity

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type Attempt struct {
	UserID      string `json:"user_id"`
	FirstCard   Card   `json:"first_card"`
	SecondCard  Card   `json:"second_card"`
	IsWin       bool   `json:"is_win"`
	Promocode   string `json:"promocode"`
	AttemptDate string `json:"attempt_date"`
}
