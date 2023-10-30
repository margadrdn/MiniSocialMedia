package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type post struct {
	Id     int    `json:"id"`
	Body   string `json:"body"`
	UserId int    `json:"user-id"`
}

var posts = []post{
	{Id: 1, Body: "Dolore aliquip ipsum ullamco enim Lorem pariatur fugiat eu fugiat.", UserId: 1},
	{Id: 2, Body: "Sit duis laboris ut laboris fugiat enim est qui amet.", UserId: 1},
	{Id: 3, Body: "Sint dolore mollit officia eu reprehenderit est minim.", UserId: 2},
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")
}
