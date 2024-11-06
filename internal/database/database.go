package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type Database struct {
	url      string
	database *sql.DB
}

func NewDatabase(databaseUrl string) *Database {
	return &Database{
		url: databaseUrl,
	}
}

// NOTE: Singleton pattern just have 1 connection
// Connection pool
func (db *Database) Connect() error {
	database, err := sql.Open(driverName, db.url)

	if err != nil {
		return err
	}

	db.database = database
	return nil
}

func (db *Database) Close() {
	db.database.Close()
}
