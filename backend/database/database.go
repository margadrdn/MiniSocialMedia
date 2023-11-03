package database

import (
	"context"
	"minisocialmedia/models"
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
		collection := client.Database("minisocialmedia").Collection("posts")

		sort := bson.D{{"CreatedAt", -1}}
		var limit int64 = 10

		cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(sort).SetLimit(limit))
		if err != nil {
			panic(err)
		}
		// end find

		var results []models.Post
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, results)
	}
}

func CreatePost(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var newPost models.Post

		if err := c.BindJSON(&newPost); err != nil {
			return
		}

		collection := client.Database("minisocialmedia").Collection("posts")

		result, err := collection.InsertOne(context.TODO(), newPost)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "error",
			})
		}

		c.IndentedJSON(http.StatusCreated, result)
	}
}
