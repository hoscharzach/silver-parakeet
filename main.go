package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var responses = []response{
	{ID: 1, Response: "It is certain."},
	{ID: 2, Response: "It is decidedly so."},
	{ID: 3, Response: "Without a doubt."},
	{ID: 4, Response: "Yes definitely."},
	{ID: 5, Response: "You may rely on it."},
	{ID: 6, Response: "As I see it, yes."},
	{ID: 7, Response: "Most likely."},
	{ID: 8, Response: "Outlook good."},
	{ID: 9, Response: "Yes."},
	{ID: 10, Response: "Signs point to yes."},
	{ID: 11, Response: "Reply hazy, try again."},
	{ID: 12, Response: "Ask again later."},
	{ID: 13, Response: "Better not tell you now."},
	{ID: 14, Response: "Cannot predict now."},
	{ID: 15, Response: "Concentrate and ask again."},
	{ID: 16, Response: "Don't count on it."},
	{ID: 17, Response: "My reply is no."},
	{ID: 18, Response: "My sources say no."},
	{ID: 19, Response: "Outlook not so good."},
	{ID: 20, Response: "Very doubtful."},
}

type indexResponse struct {
	Message string `json:"message"`
}

var message = &indexResponse{
	Message: "Hello world",
}

type response struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}

func getAllAnswers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, responses)
}

func getRandomAnswer(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	c.IndentedJSON(http.StatusOK, responses[rand.Intn(len(responses))])
}

func indexHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, message)
}

func main() {
	fmt.Println("test")

	router := gin.Default()
	router.GET("/", indexHandler)
	router.GET("/answers", getAllAnswers)
	router.GET("/random", getRandomAnswer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}