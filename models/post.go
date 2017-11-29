package models

import (
	"time"
)

const (
	PlatformHackerNews = "hacker-news"
	PlatformGoogleNews = "google-news"
	PlatformReddit     = "reddit"
	PlatformFacebook   = "facebook"
)

type Post struct {
	ID         string
	Date       time.Time
	Author     string
	ProfileImg string
	Title      string
	Content    string
	HeroImg    string
	Video      string
	IsVideo    bool
	PostLink   string
	Platform   string

	// Not sure what to do with this one
	URL string `json:"url"`

	// Reddit Specific fields
	Subreddit string  `json:"subreddit"`
	Score     int     `json:"score"`
	Created   float32 `json:"created"`
}
