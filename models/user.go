package models

type User struct {
	ID string `json:"id"`

	// TODO: Break into First and Last name?
	Name string `json:"name"`

	Username          string  `json:"username"`
	Password          string  `json:"password"`
	TwitterUsername   string  `json:"twitter-username"`
	TwitterAuthToken  string  `json:"twitter-auth-token"`
	TwitterSecret     string  `json:"twitter-secret"`
	RedditUsername    string  `json:"reddit-username"`
	RedditAuthToken   string  `json:"reddit-auth-token"`
	RedditTokenExpiry int64   `json:"reddit-token-expiry"`
	FacebookUsername  string  `json:"facebook-username"`
	FacebookAuthToken string  `json:"facebook-auth-token"`
	PostWeights       Weights `json:"post-weights"`
}
