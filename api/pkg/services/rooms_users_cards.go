package services

import (
	"cah/pkg/models"
	"context"
	"database/sql"
	"fmt"

	// Silent import
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// InsertRoomsUsersCards insert for a user a white card in hand and non used
func InsertRoomsUsersCards(RoomsusersCards *models.RoomsUsersCards) (*models.RoomsUsersCards, error) {
	sqlStatement := `
		INSERT INTO rooms_users_cards (room_id, user_id, card_id, used, on_hand)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING room_id, user_id, card_id, used, on_hand`

	RoomsusersCardsDB := &models.RoomsUsersCards{}
	err := db.QueryRow(
		sqlStatement,
		RoomsusersCards.RoomID,
		RoomsusersCards.UserID,
		RoomsusersCards.CardID,
		false,
		RoomsusersCards.OnHand).Scan(
		&RoomsusersCardsDB.RoomID,
		&RoomsusersCardsDB.CardID,
		&RoomsusersCardsDB.UserID,
		&RoomsusersCardsDB.Used,
		&RoomsusersCardsDB.OnHand)

	if err != nil {
		return nil, err
	}

	return RoomsusersCardsDB, nil
}

// UpdateRoomsUsersCards allow to update user card status
func UpdateRoomsUsersCards(RoomsusersCards *models.RoomsUsersCards) (*models.RoomsUsersCards, error) {
	sqlStatement := `
		UPDATE rooms_users_cards SET used=$1, on_hand=$2, turn=$5
		WHERE card_id=$3 AND user_id=$4
		RETURNING room_id, card_id, user_id, used, on_hand`

	RoomsusersCardsDB := &models.RoomsUsersCards{}
	err := db.QueryRow(
		sqlStatement,
		RoomsusersCards.Used,
		RoomsusersCards.OnHand,
		RoomsusersCards.CardID,
		RoomsusersCards.UserID,
		RoomsusersCards.Turn).Scan(
		&RoomsusersCardsDB.RoomID,
		&RoomsusersCardsDB.CardID,
		&RoomsusersCardsDB.UserID,
		&RoomsusersCardsDB.Used,
		&RoomsusersCardsDB.OnHand)

	if err != nil {
		return nil, err
	}

	return RoomsusersCardsDB, nil
}

// ListUserCard retrieve for a specific user in room cards in hand
func ListUserCard(userID uint64, roomID uint64) ([]*models.RoomsUsersCards, error) {
	sqlStatement := `
	SELECT
		user_id,
		card_id,
		used
	FROM
		rooms_users_cards
	WHERE
		user_id = $1
		AND room_id = $2
		AND on_hand = true
		AND used != true
	`
	rows, err := db.Query(sqlStatement, userID, roomID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	RoomsusersCards := make([]*models.RoomsUsersCards, 0)

	for rows.Next() {

		uc := &models.RoomsUsersCards{}
		if err := rows.Scan(&uc.UserID, &uc.CardID, &uc.Used); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		RoomsusersCards = append(RoomsusersCards, uc)
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

	return RoomsusersCards, nil
}

// ListUserCardDetailed retrieve for a specific user in room cards in hand (human readable)
func ListUserCardDetailed(userID uint64, roomID uint64) ([]*models.RoomsUsersCardsDetailled, error) {
	sqlStatement := `
		SELECT
			c.id,
			c.name,
			uc.user_id,
			uc.turn,
			uc.used
		FROM
			rooms_users_cards uc
		LEFT JOIN
			cards c ON c.id = uc.card_id
		LEFT JOIN
			users u ON u.id = uc.user_id
		WHERE
			user_id = $1
			AND room_id = $2
			AND on_hand = true
			AND used != true`

	rows, err := db.Query(sqlStatement, userID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	RoomsusersCardsDetailled := make([]*models.RoomsUsersCardsDetailled, 0)

	for rows.Next() {

		ucd := &models.RoomsUsersCardsDetailled{}
		if err := rows.Scan(&ucd.ID, &ucd.Name, &ucd.UserID, &ucd.Turn, &ucd.Used); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		RoomsusersCardsDetailled = append(RoomsusersCardsDetailled, ucd)
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

	return RoomsusersCardsDetailled, nil
}

// ListCurrentPlayedCards retrieve for a specific room all users cards played
func ListCurrentPlayedCards(room models.Room) ([]*models.RoomsUsersCardsDetailled, error) {
	sqlStatement := `
		SELECT
			c.id,
			c.name,
			uc.user_id,
			uc.turn,
			uc.room_id,
			uc.used
		FROM
			rooms_users_cards uc
		LEFT JOIN
			cards c ON c.id = uc.card_id
		WHERE
			uc.room_id = $1
			AND turn = $2
			AND used = true`

	rows, err := db.Query(sqlStatement, room.ID, room.Turn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	RoomsusersCardsDetailled := make([]*models.RoomsUsersCardsDetailled, 0)

	for rows.Next() {

		ucd := &models.RoomsUsersCardsDetailled{}
		if err := rows.Scan(&ucd.ID, &ucd.Name, &ucd.UserID, &ucd.Turn, &ucd.RoomID, &ucd.Used); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		RoomsusersCardsDetailled = append(RoomsusersCardsDetailled, ucd)
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

	return RoomsusersCardsDetailled, nil
}

// GetRoomsUsersCardsSimpleByCardID retrieve a specific white card in room
func GetRoomsUsersCardsSimpleByCardID(cardID uint64, roomID uint64) (*models.RoomsUsersCards, error) {
	sqlStatement := `
	SELECT
		user_id,
		card_id,
		used,
		on_hand
	FROM
		rooms_users_cards
	WHERE
		card_id = $1
		AND room_id = $2`
	uc := &models.RoomsUsersCards{}

	row := db.QueryRow(sqlStatement, cardID, roomID)
	switch err := row.Scan(&uc.UserID, &uc.CardID, &uc.Used, &uc.OnHand); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return uc, nil
	default:
		return nil, err
	}
}

// GetRoomsUsersCardsByCardID retrieve a specific white card in room (human readable)
func GetRoomsUsersCardsByCardID(cardID uint64, roomID uint64) (*models.RoomsUsersCardsDetailled, error) {
	sqlStatement := `
	SELECT
		c.ID,
		c.Name,
		uc.user_id,
		uc.turn,
		uc.room_id,
		uc.Used
	FROM
		rooms_users_cards uc
	LEFT JOIN
		cards c ON uc.card_id=c.id
	WHERE
		c.id=$1
		AND uc.room_id=$2`
	ucd := &models.RoomsUsersCardsDetailled{}

	row := db.QueryRow(sqlStatement, cardID, roomID)
	switch err := row.Scan(&ucd.ID, &ucd.Name, &ucd.UserID, &ucd.Turn, &ucd.RoomID, &ucd.Used); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return ucd, nil
	default:
		return nil, err
	}
}

// DistributeRoomsUsersCards pick new white card from room deck
func DistributeRoomsUsersCards(users []*models.User) error {
	contextLogger := logrus.WithFields(logrus.Fields{
		"model": "DistributeRoomsUsersCards",
	})
	pickCard := `
		WITH cardsUsed AS (
			SELECT
				uc.card_id AS id
			FROM
				rooms_users ru
			LEFT JOIN
				rooms_users_cards uc ON uc.room_id = ru.room_id
			WHERE
				ru.room_id = $1
				AND uc.on_hand = true
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
				c.type = 'white'
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

	giveCard := `
		INSERT INTO
			rooms_users_cards (room_id, user_id, card_id, on_hand)
		VALUES
			($1, $2, $3, $4)`

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Judge {
			continue
		}

		card := &models.Cards{}
		row := tx.QueryRowContext(ctx, pickCard, user.RoomID)
		err := row.Scan(&card.ID, &card.Description)
		if err != nil {
			contextLogger.Error("Get new cards", err)
			tx.Rollback()
			return err
		}

		contextLogger.Debugf("User %d get %d", user.ID, card.ID)

		_, err = tx.ExecContext(ctx, giveCard, user.RoomID, user.ID, card.ID, true)
		if err != nil {
			contextLogger.Error("Insert in rooms_users_cards", err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		contextLogger.Error("Commit tx", err)
		return err
	}

	return nil
}

// VodeRoomsUsersCards is well named
func VodeRoomsUsersCards(RoomsusersCards *models.RoomsUsersCards) (*models.RoomsUsersCards, error) {
	sqlStatement := `
			UPDATE rooms_users_cards SET vote = vote + 1
			WHERE card_id=$1 AND user_id=$2 AND room_id = $3
			RETURNING room_id, card_id, user_id, used, vote`

	RoomsusersCardsDB := &models.RoomsUsersCards{}
	err := db.QueryRow(
		sqlStatement,
		RoomsusersCards.CardID,
		RoomsusersCards.UserID,
		RoomsusersCards.RoomID).Scan(
		&RoomsusersCardsDB.RoomID,
		&RoomsusersCardsDB.CardID,
		&RoomsusersCardsDB.UserID,
		&RoomsusersCardsDB.Used,
		&RoomsusersCardsDB.Vote)

	if err != nil {
		return nil, err
	}

	return RoomsusersCardsDB, nil
}

// GetCountCardsByTurn is well named
func GetCountCardsByTurn(room *models.Room) (*models.CountCards, error) {
	sqlStatement := `
	SELECT
		ruc.turn,
		COUNT(ruc.card_id)
	FROM
		rooms_users_cards ruc
	WHERE
		ruc.room_id = $1
		AND ruc.turn = $2
	GROUP BY ruc.turn;`
	cc := &models.CountCards{}

	row := db.QueryRow(sqlStatement, room.ID, room.Turn)
	switch err := row.Scan(&cc.Turn, &cc.Count); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return cc, nil
	default:
		return nil, err
	}
}

// GetVotes retrieve vote by room and turn
func GetVotes(room *models.Room) (*models.CountVotes, error) {
	sqlStatement := `
	SELECT
		ruc.turn,
		SUM(ruc.vote)
	FROM
		rooms_users_cards ruc
	WHERE
		ruc.room_id = $1
		AND ruc.turn = $2
	GROUP BY ruc.turn;`
	cv := &models.CountVotes{}

	row := db.QueryRow(sqlStatement, room.ID, room.Turn)
	switch err := row.Scan(&cv.Turn, &cv.Vote); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return cv, nil
	default:
		return nil, err
	}
}

// GetElected return winner of a turn
func GetElected(room *models.Room) (*models.CardsDetailled, error) {
	sqlStatement := `
	SELECT
		c.id,
		c.name,
		ruc.user_id,
		u.name
	FROM
		rooms_users_cards ruc
	INNER JOIN
		users u	ON ruc.user_id = u.id
	INNER JOIN
		cards c	ON c.id = ruc.card_id
	WHERE
		ruc.room_id = $1
		AND ruc.turn = $2
		AND ruc.vote = 1`
	cd := &models.CardsDetailled{}

	row := db.QueryRow(sqlStatement, room.ID, room.Turn)
	switch err := row.Scan(&cd.CardID, &cd.CardName, &cd.UserID, &cd.UserName); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return cd, nil
	default:
		return nil, err
	}
}

// GetHistory return all informations cards and users by turn
func GetHistory(room models.Room) ([]models.InformationsByTurn, error) {
	sqlStatement := `
	WITH playersCards AS (
		SELECT
			c.Name as card_name,
			u.Name as user_name,
			ruc.turn
		FROM
			rooms_users_cards ruc
		INNER JOIN
			cards c ON ruc.card_id=c.id
		INNER JOIN
			users u ON u.id = ruc.user_id
		WHERE
			ruc.room_id=$1
			AND vote=1
		ORDER BY ruc.turn
	)
	SELECT
		c.name as black_card_name,
		card_name as white_card_name,
		user_name,
		pc.turn
	FROM
		playersCards pc
	INNER JOIN
		rooms_cards rc ON rc.turn = pc.turn
	INNER JOIN
		cards c ON c.id = rc.card_id
	WHERE
		rc.room_id=$1`

	rows, err := db.Query(sqlStatement, room.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	informationsByTurn := make([]models.InformationsByTurn, 0)

	for rows.Next() {
		ibt := models.InformationsByTurn{}
		if err := rows.Scan(&ibt.BlackCardName, &ibt.WhiteCardName, &ibt.UserName, &ibt.Turn); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		informationsByTurn = append(informationsByTurn, ibt)
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

	return informationsByTurn, nil
}

// GetClassementByRoom is well named
func GetClassementByRoom(room models.Room) ([]models.RoomClassement, error) {
	sqlStatement := `
	SELECT
		SUM(ruc.vote) as total,
		u.Name
	FROM
		rooms_users_cards ruc
	INNER JOIN
		users u ON u.id = ruc.user_id
	WHERE
		ruc.room_id=$1
		AND vote=1
	GROUP BY u.ID, u.Name
	ORDER BY total DESC;`

	rows, err := db.Query(sqlStatement, room.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roomClassement := make([]models.RoomClassement, 0)

	for rows.Next() {
		rc := models.RoomClassement{}
		if err := rows.Scan(&rc.Total, &rc.UserName); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		roomClassement = append(roomClassement, rc)
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

	return roomClassement, nil
}
