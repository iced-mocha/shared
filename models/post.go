package models

import (
	"time"
)

const (
	PlatformHackerNews = "Hacker News"
)

type Post struct {
	ID         string
	Date       time.Time
	Author     string
	ProfileImg string
	Title      string
	Content    string
	HeroImg    string
	PostLink   string
	Platform   string
}
