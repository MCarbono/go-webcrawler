package entities

import (
	"regexp"
	"time"
)

type Feed struct {
	Title       string
	Description string
	Provider    string
	Link        string
	Date        string
	CreatedAt   time.Time
}

func NewFeed(title, description, link, date string) *Feed {
	f := &Feed{
		Title:       title,
		Description: description,
		Link:        link,
		Date:        date,
		Provider:    link,
		CreatedAt:   time.Now(),
	}
	r, err := regexp.Compile(`https://(w{3}\.)?`)
	if err != nil {
		return f
	}
	s := r.ReplaceAllString(link, "")
	r, err = regexp.Compile(`\..*`)
	if err != nil {
		return f
	}
	f.Provider = r.ReplaceAllString(s, "")
	return f
}
