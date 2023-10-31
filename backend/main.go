package main

import (
	"context"
	"log"
	"net/http"

	"minisocialmedia/environment"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// Get environment value
	uri, err := environment.GetEnvVar("MONGODB_URI")
	if err != nil {
		log.Fatal(err)
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

	coll := client.Database("minisocialmedia").Collection("posts")
	filter := bson.D{{"author", "spider_man"}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// end find

	var results []post
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	posts = results

	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")
}
