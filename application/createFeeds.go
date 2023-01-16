package application

import (
	"go-webscrapper/domain/entities"
	"go-webscrapper/domain/repository"
	"log"
	"sync"

	"github.com/mmcdole/gofeed"
)

type CreateFeeds struct {
	repository.FeedRepository
	Scrapper *gofeed.Parser
}

func (uc CreateFeeds) Execute(input FeedInput) {
	var wg sync.WaitGroup
	var ch = make(chan *entities.Feed)
	go saveFeed(ch, uc.FeedRepository)
	createScrappers(input, &wg, ch, uc.Scrapper)
	wg.Wait()
	close(ch)
}

func createScrappers(input FeedInput, wg *sync.WaitGroup, ch chan<- *entities.Feed, scrapper *gofeed.Parser) {
	for _, v := range input {
		wg.Add(1)
		go scrappWebPage(v, wg, ch, scrapper)
	}
}

func scrappWebPage(link string, wg *sync.WaitGroup, ch chan<- *entities.Feed, scrapper *gofeed.Parser) {
	defer wg.Done()
	feed, err := scrapper.ParseURL(link)
	if err != nil {
		log.Printf("Some error occurred while trying to scrapp the %v page. This website will be ignored. Err: %v", link, err)
		return
	}
	for _, v := range feed.Items {
		ch <- entities.NewFeed(v.Title, v.Description, v.Link, v.Published)
	}
}

func saveFeed(ch <-chan *entities.Feed, repo repository.FeedRepository) {
	for msg := range ch {
		if err := repo.Create(msg); err != nil {
			log.Printf("[Create Feed] error: %v\n Feed %+v", err, msg)
		}
	}
}

type FeedInput []string
