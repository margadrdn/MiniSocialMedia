package main

import (
	"minisocialmedia/authentication"
	"minisocialmedia/database"
	"minisocialmedia/environment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get environment value
	uri, err := environment.GetEnvVar("MONGODB_URI")
	if err != nil {
		panic(err)
	}

	// Connect to database
	client, err := database.Connect(uri)
	if err != nil {
		panic(err)
	}

	// Disconnect from database after server shutdown
	defer func() {
		if err := database.Disconnect(client); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	router.GET("/posts", database.GetPosts(client))
	router.POST("/posts", database.CreatePost(client))
	router.POST("/auth/login", authentication.SigninUser(client)) // todo change function name
	router.POST("/auth/signup", authentication.SignupUser(client))
	router.Run("localhost:8080")
}
