package authentication

import (
	"context"
	"fmt"
	"minisocialmedia/environment"
	"minisocialmedia/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SigninUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}

		collection := client.Database("minisocialmedia").Collection("users")
		filter := bson.D{{"name", user.Name}}

		var databaseUser models.User

		err := collection.FindOne(context.TODO(), filter).Decode(&databaseUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "wrong username"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(user.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "wrong password"})
			return
		}

		token, err := createToken(databaseUser.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func SignupUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}

		collection := client.Database("minisocialmedia").Collection("users")
		filter := bson.D{{"name", user.Name}}
		// var databaseUser models.User
		count, _ := collection.CountDocuments(context.TODO(), filter)
		if count == 0 {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err,
				})
			}
			user.Password = string(hash)
			result, err := collection.InsertOne(context.TODO(), user)

			fmt.Println(result)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err,
				})
			}
			jwt, err := createToken(user.Name)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			if jwt != "" {
				c.JSON(http.StatusCreated, jwt)
			}
		} else if count != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "username already exists",
			})
		}
	}
}

type UserClaims struct {
	username string
	jwt.StandardClaims
}

func createToken(username string) (string, error) {
	secret, err := environment.GetEnvVar("SECRET")
	if err != nil {
		return "", err
	}
	claims := UserClaims{
		username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
