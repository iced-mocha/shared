package models

type User struct {
	ID string `json:"id"`

	// TODO: Break into First and Last name?
	Name string `json:"name"`

	Username   string `json:"username"`
	Password   string `json:"password"`
	RedditUser string `json:"reddit-user"`
}
