package models

type Post struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
