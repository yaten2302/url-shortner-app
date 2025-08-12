package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type routeURL struct {
	Url string `form:"url"`
}

type URL struct {
	OrignalURL    string `bson:"original_url"`
	CompressedURL string `bson:"compressed_url"`
}

var client *mongo.Client

func main() {
	router := gin.Default()
	router.GET("/shorten", shortenRoute)

	client = mongoDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	err := router.Run()

	if err != nil {
		log.Println("err")
	}
}

func shortenRoute(c *gin.Context) {
	var url routeURL

	if err := c.ShouldBind(&url); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	shortenedURL := GenerateShortenedURL()

	coll := client.Database("url-shortener-app").Collection("URLs")
	doc := URL{
		OrignalURL:    url.Url,
		CompressedURL: shortenedURL,
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to insert"})
		return
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	c.String(200, "http://localhost:8080/"+shortenedURL)
}

func GenerateShortenedURL() string {
	var (
		randomChars   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
		randIntLength = 27
		stringLength  = 32
	)

	str := make([]rune, stringLength)

	for char := range str {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(randIntLength)))
		if err != nil {
			panic(err)
		}

		str[char] = randomChars[nBig.Int64()]
	}

	hash := sha256.Sum256([]byte(uniqid(string(str))))
	encodedString := base64.StdEncoding.EncodeToString(hash[:])

	return encodedString[0:9]
}

func uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000

	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}
