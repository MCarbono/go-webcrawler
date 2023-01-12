package main

import (
	"fmt"
	"go-webscrapper/application"
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
	startAt := time.Now()
	createFeeds.Execute([]string{
		"https://catracalivre.com.br/feed/",
		"https://www.infomoney.com.br/feed/",
		"https://forbes.com.br/feed/",
		"https://www.cnnbrasil.com.br/feed/",
		"https://www.moneytimes.com.br/feed/",
	})
	endAt := time.Now()
	diff := endAt.Sub(startAt)
	fmt.Printf("Pages has been scrapped. \nTime elapsed: %v", diff.Seconds())
}
