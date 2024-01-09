package main

import (
	"fmt"
    "net/http"
	"github.com/gin-gonic/gin"
)

type server struct {
	port string
	router *gin.Engine
}

// server methods
func (server *server) Init(port string) {
	server.port = port

	server.router = gin.Default()
	server.router.GET("/ping", ping)
	server.router.GET("/", home)

	// All routes defined within this group will have the /v1 prefix
	apiV1 := server.router.Group("/v1")
	apiV1.POST("/shorten", routes.ShortenURL)
	apiV1.GET("/url", routes.GetLongURL)
}

func (server *server) Serve() error {
	return server.router.Run(fmt.Sprintf(":%v", server.port))
}

// router functions
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "reached home"})
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong"
	})
}

// =========NOTES============
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
