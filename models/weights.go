package models

// Used for unmarshaling weightings for different clients
type Weights struct {
	Reddit     float64            `json:"reddit"`
	Facebook   float64            `json:"facebook"`
	HackerNews float64            `json:"hacker-news"`
	GoogleNews float64            `json:"google-news"`
	Twitter    float64            `json:"twitter"`
	RSS        map[string]float64 `json:"rss"`
}
