package services

import (
	"cah/pkg/models"

	// pq import directive
	_ "github.com/lib/pq"
)

// InsertRoomsUsers is well named
func InsertRoomsUsers(user *models.User) (*models.RoomCards, error) {
	sqlStatement := `
		INSERT INTO rooms_users (user_id, room_id)
		VALUES ($1, $2)
		RETURNING user_id, room_id`

	roomCardsDB := &models.RoomCards{}
	err := db.QueryRow(
		sqlStatement,
		user.ID,
		user.RoomID).Scan(
		&roomCardsDB.UserID,
		&roomCardsDB.RoomID)

	if err != nil {
		return nil, err
	}

	return roomCardsDB, nil
}
