package models

type Post struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
