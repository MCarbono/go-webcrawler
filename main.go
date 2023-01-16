package main

import (
	"fmt"
	"go-webscrapper/application"
	"go-webscrapper/infra/controller"
	"go-webscrapper/infra/database"
	"go-webscrapper/infra/repository"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mmcdole/gofeed"
)

func main() {
	db, err := database.NewConnection("sqlite3", "./go-webscrapper.db")
	if err != nil {
		panic(err)
	}
	feedRepository := repository.NewFeedRepositorySQL(db)
	createFeeds := application.CreateFeeds{
		FeedRepository: feedRepository,
		Scrapper:       gofeed.NewParser(),
	}
	feedController := controller.NewFeedController(createFeeds)
	startAt := time.Now()
	feedController.Execute()
	endAt := time.Now()
	diff := endAt.Sub(startAt)
	fmt.Printf("Pages has been scrapped. \nTime elapsed: %v", diff.Seconds())
}
