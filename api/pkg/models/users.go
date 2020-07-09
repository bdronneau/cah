package models

// User is the DB representation
type User struct {
	ID          uint64 `json:"id" form:"id" query:"id"`
	Name        string `json:"name" form:"name" query:"name"`
	RoomName    string `json:"room_name" form:"room_name" query:"room_name"`
	RoomID      uint64 `json:"room_id" form:"room_id" query:"room_id"`
	Judge       bool   `json:"judge"`
	LastUpdated string `json:"lastupdated" form:"lastupdated" query:"lastupdated"`
}

// UserPostCard is used for HTTP POST
type UserPostCard struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Turn   uint64 `json:"turn"`
}
