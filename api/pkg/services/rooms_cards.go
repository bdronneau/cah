package services

import (
	"database/sql"
	"fmt"

	"cah/pkg/models"

	// pq import directive
	_ "github.com/lib/pq"
)

// InsertRoomsCards put by with turn a black card
func InsertRoomsCards(roomCard *models.RoomsCards) (*models.RoomsCards, error) {
	sqlStatement := `
		INSERT INTO rooms_cards (card_id, room_id, used, turn)
		VALUES ($1, $2, $3, $4)
		RETURNING card_id, room_id, used, turn`

	roomCardDB := &models.RoomsCards{}
	err := db.QueryRow(
		sqlStatement,
		roomCard.CardID,
		roomCard.RoomID,
		false,
		roomCard.Turn).Scan(
		&roomCardDB.CardID,
		&roomCardDB.RoomID,
		&roomCardDB.Used,
		&roomCardDB.Turn)

	if err != nil {
		return nil, err
	}

	return roomCardDB, nil
}

// GetRoomsCardsByTurn retrieve for specific roomID and turn the black card
func GetRoomsCardsByTurn(room *models.Room) (*models.RoomsCardsDetailled, error) {
	sqlStatement := `SELECT c.ID , c.Name, rc.Turn FROM rooms_cards rc LEFT JOIN cards c ON rc.card_id=c.id WHERE rc.room_id=$1 AND rc.turn=$2`
	rcd := &models.RoomsCardsDetailled{}

	row := db.QueryRow(sqlStatement, room.ID, room.Turn)
	switch err := row.Scan(&rcd.ID, &rcd.Name, &rcd.Turn); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return rcd, nil
	default:
		return nil, err
	}
}
