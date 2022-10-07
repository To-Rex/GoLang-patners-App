package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


const uri = "mongodb+srv://root:1234@cluster0.ik76ncs.mongodb.net/?retryWrites=true&w=majority"


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Token	string `json:"token"`
}

type Token struct {
	Token string `json:"token"`
}

func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.POST("/register", register)
	r.GET("/getuser", user)
	//r.GET("/getusers", users)
	//r.PUT("/updateuser", updateuser)
	r.Run()
}

func login(c *gin.Context) {
	var user User
	var token Token

	c.BindJSON(&user)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("test").Collection("users")

	filter := bson.M{"username": user.Username, "password": user.Password}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token.Token = createToken(user.Username)
	c.JSON(http.StatusOK, token)
}

func register(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user.Token = createToken(user.Username)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("test").Collection("users")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,"Token": user.Token})
}

func user(c *gin.Context) {
	var user User
	var token Token
	c.BindJSON(&token)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("test").Collection("users")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"token": token.Token}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, user)
}


func createToken(username string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString
}