package models

// Room is the DB representation
type Room struct {
	ID          uint64 `json:"id" form:"id" query:"id"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	Turn        uint64 `json:"turn" form:"turn" query:"turn"`
	Status      string `json:"status" form:"status" query:"status"`
	LastUpdated string `json:"lastupdated" form:"lastupdated" query:"lastupdated"`
}
