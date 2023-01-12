package database

import "database/sql"

func NewConnection(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	err = seedDatabase(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func seedDatabase(db *sql.DB) error {
	_, err := db.Exec("drop table if exists feeds;")
	if err != nil {
		return err
	}
	_, err = db.Exec("create table feeds (title text, description text, provider text, link text unique, date text, created_at date);")
	return err
}
