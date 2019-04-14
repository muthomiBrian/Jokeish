package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Joke contains information about a single joke
type Joke struct {
	ID	int 	`json:"id" binding:"required"`
	Likes int	`json:"likes"`
	Joke	string	`json:"joke" binding:"required"`
}

// Create a list of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmospher."},
	Joke{2, 0, "What do you call a fake noodle? An impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable"},
	Joke{5, 0, "I just watched a program about beavers, It was the best dam program I've ever seen"},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {
	// Set the router as the default shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message":"pong",
			})
		})
	}

	
	// Our API will consist of two routes
	// /jokes - will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - will handle likes sent to a joke
	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	// Start and run the server
	router.Run(":3000")
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK,jokes)
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	// Confirm joke ID sent is valid
	// remember to import the `strconv` package
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for index := range jokes {
			if jokes[index].ID == jokeid {
				jokes[index].Likes ++
			}
		}
	} else {
		// Joke ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}