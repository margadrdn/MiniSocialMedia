package main

import (
	"minisocialmedia/database"
	"minisocialmedia/environment"

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
	router.GET("/posts", database.GetPosts(client))
	router.Run("localhost:8080")
}
