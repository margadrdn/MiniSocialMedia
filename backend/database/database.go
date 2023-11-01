package database

import (
	"context"
	"minisocialmedia/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (client *mongo.Client, err error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return
}

func Disconnect(client *mongo.Client) (err error) {
	err = client.Disconnect(context.TODO())
	return
}

func GetPosts(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		coll := client.Database("minisocialmedia").Collection("posts")
		filter := bson.D{{"author", "spider_man"}}

		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		// end find

		var results []model.Post
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, results)
	}
}
