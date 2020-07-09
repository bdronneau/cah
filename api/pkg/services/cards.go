package services

import (
	"database/sql"
	"fmt"

	"cah/pkg/models"
)

// ListCardsByType retrieve cards by type
func ListCardsByType(typeCard string) ([]*models.Cards, error) {
	rows, err := db.Query("SELECT id, name FROM cards WHERE type=$1", typeCard)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]*models.Cards, 0)

	for rows.Next() {

		c := &models.Cards{}
		if err := rows.Scan(&c.ID, &c.Description); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		cards = append(cards, c)
	}

	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()

	if rerr != nil {
		return nil, rerr
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

// PickNewBlackCard pick randomly a new black card. Played black card are in deck.
func PickNewBlackCard(room models.Room) (models.Cards, error) {
	sqlStatement := `
	WITH cardsUsed AS (
		SELECT
			rc.card_id AS id
		FROM
			rooms_cards rc
		WHERE
			rc.room_id = $1
	), randomDeck AS (
		SELECT
			c.id,
			c.name,
			c.type,
			random() as ordering
		FROM
			cards c
		LEFT JOIN
			cardsUsed cu ON cu.id = c.id
		WHERE
			c.type = 'black'
			AND cu.id IS NULL
		ORDER BY
			ordering
	)
	SELECT
		id,
		name
	FROM
		randomDeck
	LIMIT 1;`
	c := models.Cards{}

	switch err := db.QueryRow(sqlStatement, room.ID).Scan(&c.ID, &c.Description); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return c, err
	case nil:
		return c, nil
	default:
		return c, err
	}
}
