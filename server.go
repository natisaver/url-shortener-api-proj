package main

import (
    "net/http"
	"github.com/gin-gonic/gin"
	// "github.com/natisaver/urlshortner/routes"
	"github.com/natisaver/url-shortener-api-proj/urlshortner/routes"
)

type server struct {
	port string
	router *gin.Engine
}

// server methods
func (s *server) Init(port string) {
	s.port = port
	s.router = gin.Default()

	s.router.GET("/ping", ping)
	s.router.GET("/", home)

	// All routes defined within this group will have the /v1 prefix
	apiV1 := s.router.Group("/v1")
	apiV1.POST("/shorten", routes.ShortenURL)
	apiV1.GET("/:encodedurl", routes.GetLongURL)
}

func (s *server) Serve() error {
	return s.router.Run("localhost:" + s.port)
}

// router functions
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "reached home"})
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// =========NOTES ON GOLANG METHODS============
// format of methods
// func (obj MyType) MethodName() (returnType) {}

// Value Receiver
// does not affect original instance
// type Circle struct {
//     radius float64
// }

// func (c Circle) Area() float64 {
//     return 3.14 * c.radius * c.radius
// }

// Pointer Receiver
// affects the original instance
// type Counter struct {
//     count int
// }

// func (c *Counter) Increment() {
//     c.count++
// }
