package services

import (
	"database/sql"
	"fmt"
	"time"

	"cah/pkg/models"
	"cah/pkg/utils"
)

// ListRoom is well named
func ListRoom() ([]*models.Room, error) {
	rows, err := db.Query("SELECT id, name, description, status, turn FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]*models.Room, 0)

	for rows.Next() {
		room := &models.Room{}
		if err := rows.Scan(&room.ID, &room.Name, &room.Description, &room.Status, &room.Turn); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return nil, err
		}

		rooms = append(rooms, room)
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

	return rooms, nil
}

// InsertRoom is well named
func InsertRoom(description string) (*models.Room, error) {
	sqlStatement := `
		INSERT INTO rooms (name, description, turn, status, lastupdated)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, turn, status`
	roomDB := &models.Room{}
	timeUpdated, _ := timeIn(time.Now(), "")

	err := db.QueryRow(
		sqlStatement,
		utils.RandSeq(),
		description,
		0,
		"created",
		timeUpdated).Scan(
		&roomDB.ID,
		&roomDB.Name,
		&roomDB.Turn,
		&roomDB.Status)
	if err != nil {
		return nil, err
	}

	return roomDB, nil
}

// GetRoomName is well named
func GetRoomName(roomName string) (*models.Room, error) {
	sqlStatement := `SELECT id, name, description, turn, status FROM rooms WHERE name=$1;`
	room := &models.Room{}

	row := db.QueryRow(sqlStatement, roomName)
	switch err := row.Scan(&room.ID, &room.Name, &room.Description, &room.Turn, &room.Status); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, err
	case nil:
		return room, nil
	default:
		return nil, err
	}
}

// TODO: move to helpers
func timeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// UpdateRoom is well named
func UpdateRoom(room *models.Room) (*models.Room, error) {
	sqlStatement := `
		UPDATE rooms SET description=$1, turn=$2, status=$3, lastupdated=$4 WHERE id=$5
		RETURNING id, name, turn, status`
	roomDB := &models.Room{}
	timeUpdated, _ := timeIn(time.Now(), "")

	err := db.QueryRow(
		sqlStatement,
		room.Description,
		room.Turn,
		room.Status,
		timeUpdated,
		room.ID).Scan(
		&roomDB.ID,
		&roomDB.Name,
		&roomDB.Turn,
		&roomDB.Status)
	if err != nil {
		return nil, err
	}

	return roomDB, nil
}
