package routes

import (
github.com/api

	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/natisaver/url-shortener-api-proj/urlshortner/utils"
	"github.com/natisaver/url-shortener-api-proj/urlshortner/models/url"
)

func ShortenURL(c *gin.Context) {
	// Parse the JSON request body into a URL object
	var request url.URL
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// sanitise for security
	sanitizedURL, err := utils.SanitizeURL(request.longURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// shorten it to an encoded form
	shortenedURL := utils.ShortenURL(sanitizedURL)

	// Create a new object with the modified shortURL
    modifiedRequest := url.URL{
        LongURL:   sanitizedURL,
        ShortURL:  shortenedURL,
		CreatedAt: time.Now(),
    }

	// store into db
	// we first create a transaction i.e. a temporary form of our DB connection
	tx := common.GetDB().Begin()
	// instead of passing in the actual DB connection
	// this is to ensure all queries are in one transaction, ensuring consistency of data
	err := url.StoreURL(tx, request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	c.JSON(http.StatusOK, modifiedRequest)
}

func GetLongURL(c *gin.Context) {
	shortenedurl := c.Param("encodedurl")

}

// =======NOTES ON BINDINGJSON========
// Use ShouldBindJSON if you want to handle binding errors yourself and continue with the request flow.
// Use BindJSON if you prefer a simpler way to handle binding errors by responding with a 400 Bad Request status directly.