package models

// RoomsUsersCards is the DB representation
type RoomsUsersCards struct {
	RoomID uint64 `json:"room_id"`
	CardID uint64 `json:"card_id"`
	UserID uint64 `json:"user_id"`
	Used   bool   `json:"used"`
	Vote   int    `json:"vote"`
	OnHand bool   `json:"on_hand"`
	Turn   uint64 `json:"turn"`
}

// InformationsByTurn for a room
type InformationsByTurn struct {
	BlackCardName string `json:"black_card_name"`
	WhiteCardName string `json:"white_card_name"`
	UserName      string `json:"user_name"`
	Turn          uint64 `json:"turn"`
}

// RoomClassement for the room resume
type RoomClassement struct {
	Total    uint64 `json:"total"`
	UserName string `json:"user_name"`
}

// RoomsUsersCardsDetailled is the human readable user card
type RoomsUsersCardsDetailled struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	UserID uint64 `json:"user_id"`
	RoomID uint64 `json:"room_id"`
	Turn   uint64 `json:"turn"`
	Used   bool   `json:"used"`
}

// CardsDetailled is the human readable user card
type CardsDetailled struct {
	CardID   string `json:"card_id"`
	CardName string `json:"card_name"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// CountCards of turn
type CountCards struct {
	Turn  uint64 `json:"turn"`
	Count int    `json:"count"`
}

// CountVotes of turn
type CountVotes struct {
	Turn uint64 `json:"turn"`
	Vote int    `json:"vote"`
}
