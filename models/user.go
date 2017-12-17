package models

type User struct {
	ID string `json:"id"`

	// TODO: Break into First and Last name?
	Name string `json:"name"`

	Username           string              `json:"username"`
	Password           string              `json:"password"`
	TwitterUsername    string              `json:"twitter-username"`
	TwitterAuthToken   string              `json:"twitter-auth-token"`
	TwitterSecret      string              `json:"twitter-secret"`
	RedditUsername     string              `json:"reddit-username"`
	RedditAuthToken    string              `json:"reddit-auth-token"`
	RedditRefreshToken string              `json:"reddit-refresh-token"`
	RedditTokenExpiry  int64               `json:"reddit-token-expiry"`
	FacebookUsername   string              `json:"facebook-username"`
	FacebookAuthToken  string              `json:"facebook-auth-token"`
	RssGroups          map[string][]string `json:"rss-groups"`
	PostWeights        Weights             `json:"post-weights"`
}
