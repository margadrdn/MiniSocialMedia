package model

type Post struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}
