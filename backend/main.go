package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI fatal error. uri not set correctly")
	}
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")
}
