package models

import (
	"time"
)

const (
	PlatformHackerNews = "hacker-news"
	PlatformGoogleNews = "google-news"
	PlatformReddit     = "reddit"
	PlatformFacebook   = "facebook"
	PlatformRSS        = "rss"
)

type Post struct {
	ID          string
	Author      string
	DisplayName string
	ProfileImg  string
	Title       string
	Content     string
	HeroImg     string
	Video       string
	PostLink    string
	Platform    string
	IsVideo     bool
	Date        time.Time
	Retweets    int `json:"retweets"`
	Favourites  int `json:"favourites"`

	// Not sure what to do with this one
	URL  string `json:"url"`
	Meta string `json:"meta"`

	// Reddit Specific fields
	Subreddit string  `json:"subreddit"`
	Score     int     `json:"score"`
	Created   float32 `json:"created"`
}
