package db

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	// blank import for postgres driver
	_ "github.com/lib/pq"
)

// Adapter implements the DbPort interface
type Adapter struct {
	db *sql.DB
}

// NewAdapter creates a new Adapter
func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

// CloseDbConnection closes the db  connection
func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

// Company
// Insert new company entry to the database companies table
func (da Adapter) InsertNewCompanyToDb(uuid string, name string, contact string, website string, email string, country string, state string, zip string) (int32, error) {
	queryString, args, err := sq.Insert("companies").Columns("uuid", "name", "contact", "website", "email", "country", "state", "zip").
		Values(uuid, name, contact, website, email, country, state, zip).ToSql()

	if err != nil {
		return 0, err
	}

	_, err = da.db.Exec(queryString, args...)

	if err != nil {
		return 0, err
	}

	return 0, nil
}