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
	Retweets    int

	// Not sure what to do with this one
	URL string `json:"url"`

	// Reddit Specific fields
	Subreddit string  `json:"subreddit"`
	Score     int     `json:"score"`
	Created   float32 `json:"created"`
}
