package repository

import (
	"database/sql"
	"go-webscrapper/domain/entities"
)

type FeedRepositorySQL struct {
	connection *sql.DB
}

func NewFeedRepositorySQL(connection *sql.DB) *FeedRepositorySQL {
	return &FeedRepositorySQL{
		connection: connection,
	}
}

func (r FeedRepositorySQL) Create(e *entities.Feed) error {
	stmt, err := r.connection.Prepare("INSERT OR IGNORE INTO feeds (title, description, provider, link, date, created_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(e.Title, e.Description, e.Provider, e.Link, e.Date, e.CreatedAt)
	return err
}
