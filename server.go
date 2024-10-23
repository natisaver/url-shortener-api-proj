package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/natisaver/urlshortner/routes"
	"urlshortener/handlers"
)

type Server struct {
	port   string
	router *gin.Engine
}

// Server methods
func (s *Server) Init(port string) {
	s.port = port
	s.router = gin.Default()

	// add cors middleware as frontend and server are on different ports
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add frontend URL
	s.router.Use(cors.New(config))

	s.router.GET("/ping", ping)
	s.router.GET("/", home)

	// All routes defined within this group will have the /v1 prefix
	apiV1 := s.router.Group("/v1")
	apiV1.POST("/shorten", handlers.ShortenURL)
	apiV1.GET("/:encodedurl", handlers.GetLongURL)
}

func (s *Server) Serve() error {
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
