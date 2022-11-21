package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

const uri = "mongodb+srv://root:1234@cluster0.ik76ncs.mongodb.net/?retryWrites=true&w=majority"

type User struct {
	Name	 string `json:"name"`
	Surname  string `json:"surname"`
	Years	 int	`json:"years"`
	Phone	 string `json:"phone"`
	Country string `json:"country"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	RegisterDate string `json:"register_date"`
	Money   int	`json:"money"`
	Promocode string `json:"promocode"`
	Verified bool `json:"verified"`
	Blocked bool `json:"blocked"`
	Token string `json:"token"`
	UserId string `json:"user_id"`
	UserStatus string `json:"user_status"`
	UserRole string `json:"user_role"`
	UserAvatar string `json:"user_avatar"`
	Wallet string `json:"wallet"`
}

type Transaction struct {
	TransactionId string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	TransactionType string `json:"transaction_type"`
	TransactionAmount int `json:"transaction_amount"`
	TransactionStatus string `json:"transaction_status"`
	TransactionCurrency string `json:"transaction_currency"`
	TransactionDescription string `json:"transaction_description"`
	TransactionUserId string `json:"transaction_user_id"`
}

type Product struct {
	ProductId string `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice int `json:"product_price"`
	ProductCurrency string `json:"product_currency"`
	ProductImage string `json:"product_image"`
	ProductCategory string `json:"product_category"`
	ProductStatus string `json:"product_status"`
	ProductUserId string `json:"product_user_id"`
}

type Order struct {
	OrderId string `json:"order_id"`
	OrderDate string `json:"order_date"`
	OrderStatus string `json:"order_status"`
	OrderUserId string `json:"order_user_id"`
	OrderProductId string `json:"order_product_id"`
	OrderProductAmount int `json:"order_product_amount"`
	OrderProductPrice int `json:"order_product_price"`
	OrderProductCurrency string `json:"order_product_currency"`
}

func main() {
	router := gin.Default()
	router.POST("/register", register)
	router.Run(":8080")
}

func createToken(username string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString
}

func passwordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

//generate userid random cheracter 32 length string
func generateUserId() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 32
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

//generate wallet random cheracter 32 length string for user
func generateWallet() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 32
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Name	 string `json:"name"`
	Surname  string `json:"surname"`
	Years	 int	`json:"years"`
	Phone	 string `json:"phone"`
	Country string `json:"country"`
	Email	 string `json:"email"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	var register Register
	c.BindJSON(&register)
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
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("Partners").Collection("users")
	var user User
	err = collection.FindOne(ctx, bson.M{"email": register.Email}).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	if user.Email == register.Email {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "User already exists"})
		return
	}
	hash := passwordHash(register.Password)
	userId := generateUserId()
	wallet := generateWallet()
	user = User{
		Name: register.Name,
		Surname: register.Surname,
		Years: register.Years,
		Phone: register.Phone,
		Country: register.Country,
		Email: register.Email,
		Password: hash,
		RegisterDate: time.Now().Format("2006-01-02 15:04:05"),
		Money: 0,
		Promocode: "",
		Verified: false,
		Blocked: false,
		Token: createToken(register.Email),
		UserId: userId,
		UserStatus: "user",
		UserRole: "user",
		UserAvatar: "",
		Wallet: wallet,
	}
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User created successfully", "data": user})
	return 
}
