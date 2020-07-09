package models

// RoomCards is the DB representation
type RoomCards struct {
	UserID uint64 `json:"user_id" form:"user_id" query:"user_id"`
	RoomID uint64 `json:"room_id" form:"room_id" query:"room_id"`
}
