package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"cah/pkg/models"

	"github.com/sirupsen/logrus"
)

// InsertUser is well named
func InsertUser(user *models.User) (*models.User, error) {
	sqlStatement := `
		INSERT INTO users (name, lastupdated)
		VALUES ($1, $2)
		RETURNING id, name`

	userDB := &models.User{}
	timeUpdated, _ := timeIn(time.Now(), "")
	err := db.QueryRow(
		sqlStatement,
		user.Name,
		timeUpdated).Scan(
		&userDB.ID,
		&userDB.Name)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}

// ListUserByRoom is well named
func ListUserByRoom(roomID uint64) ([]*models.User, error) {
	sqlStatement := `
		SELECT
			u.id,
			u.name,
			r.id,
			ur.judge
		FROM
			rooms r
		LEFT JOIN
			rooms_users ur ON ur.room_id = r.id
		LEFT JOIN
			users u ON u.id = ur.user_id
		WHERE
			r.id = $1
			AND ur.room_id IS NOT NULL`

	rows, err := db.Query(sqlStatement, roomID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {

		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.RoomID, &user.Judge); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		users = append(users, user)
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

	return users, nil
}

// UserElected move judge iin transaction
func UserElected(user models.User) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// TODO: use select to get current judge
	_, err = tx.ExecContext(ctx, `UPDATE rooms_users SET judge=false WHERE room_id=$1`, user.RoomID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE rooms_users SET judge=true WHERE room_id=$1 AND user_id=$2`, user.RoomID, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetJudge is well named
func GetJudge(room *models.Room) (*models.User, error) {
	sqlStatement := `SELECT u.id, name, room_id, judge FROM rooms_users rc LEFT JOIN users u ON u.id = rc.user_id WHERE rc.room_id = $1 AND rc.judge = true`
	user := &models.User{}

	row := db.QueryRow(sqlStatement, room.ID)
	switch err := row.Scan(&user.ID, &user.Name, &user.RoomID, &user.Judge); err {
	case sql.ErrNoRows:
		logrus.Error("No rows were returned")
		return nil, err
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// SetNextJudge is well named
func SetNextJudge(room models.Room) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE rooms_users SET judge=false WHERE room_id=$1 AND judge=true`, room.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	var newJudgeID uint64
	stmtCurrentWinner := `
	SELECT
		user_id
	FROM
		rooms_users_cards ruc
	WHERE
		ruc.room_id = $1
		AND ruc.vote = 1
		AND ruc.turn = $2`

	err = tx.QueryRowContext(ctx, stmtCurrentWinner, room.ID, room.Turn).Scan(&newJudgeID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE rooms_users SET judge=true WHERE room_id=$1 AND user_id=$2`, room.ID, newJudgeID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetUser is well named
func GetUser(id uint64) (models.User, error) {
	sqlStatement := `SELECT id, name FROM users WHERE id=$1;`
	user := models.User{}

	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&user.ID, &user.Name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, err
	case nil:
		return user, nil
	default:
		return user, err
	}
}
