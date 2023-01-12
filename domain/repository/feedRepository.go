package repository

import "go-webscrapper/domain/entities"

type FeedRepository interface {
	Create(feed *entities.Feed) error
}
