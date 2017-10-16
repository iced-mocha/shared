package models

import ()

type ClientResp struct {
	Posts   []Post
	NextURL string
}
