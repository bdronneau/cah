package models

// RoomsCards is the link for the black cards
type RoomsCards struct {
	CardID uint64 `json:"card_id" form:"card_id" query:"card_id"`
	RoomID uint64 `json:"room_id" form:"room_id" query:"room_id"`
	Used   bool   `json:"used" form:"used" query:"used"`
	Turn   uint64 `json:"turn" form:"turn" query:"turn"`
}

// RoomsCardsDetailled is the representation of the black card
type RoomsCardsDetailled struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Turn uint64 `json:"turn"`
}
