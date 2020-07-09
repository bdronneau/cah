package services

import (
	"database/sql"

	// pq import directive
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

// InitDB initialize connector to postgres database
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		logrus.Panicf("Trouble on connect to PG %v", err)
	}

	if err = db.Ping(); err != nil {
		logrus.Panicf("Trouble on PG ping %v", err)
	}
}
