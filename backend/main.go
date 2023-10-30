package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type post struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

var posts = []post{
	{Id: uuid.NewString(), Author: "margad", Body: "Dolore aliquip ipsum ullamco enim Lorem pariatur fugiat eu fugiat."},
	{Id: uuid.NewString(), Author: "catperson", Body: "Sit duis laboris ut laboris fugiat enim est qui amet."},
	{Id: uuid.NewString(), Author: "spider-man", Body: "Sint dolore mollit officia eu reprehenderit est minim."},
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")
}
