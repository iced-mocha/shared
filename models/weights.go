package models

// Used for unmarshaling weightings for different clients
type Weights struct {
	Reddit     int `json:"reddit"`
	Facebook   int `json:"facebook"`
	HackerNews int `json:"hacker-news"`
	GoogleNews int `json:"google-news"`
}
