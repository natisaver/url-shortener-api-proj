package routes

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"urlshortener/common"
	"urlshortener/models/url"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ShortenURL(c *gin.Context) {
	// Parse the JSON request body into a URL object
	var request url.URL
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("Error parsing JSON request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Sanitize for security
	sanitizedURL, err := utils.SanitizeURL(request.LongURL)
	if err != nil {
		fmt.Println("Error sanitizing URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL"})
		return
	}

	// Shorten the URL to an encoded form
	shortenedURL, err := utils.ShortenURL(sanitizedURL)
	if err != nil {
		fmt.Println("Error shortening URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid URL"})
		return
	}

	// Create a new object with the modified shortURL
	modifiedRequest := url.URL{
		LongURL:   sanitizedURL,
		ShortURL:  shortenedURL,
		CreatedAt: time.Now(),
	}

	// Open the database
	db, err := common.GetDB()
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		return
	}
	defer db.Close()

	// Store into the database
	// We first create a transaction, i.e., a temporary form of our DB connection
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error beginning database transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		return
	}

	// Instead of passing in the actual DB connection,
	// this is to ensure all queries are in one transaction, ensuring consistency of data
	err = url.StoreURLWithTransaction(tx, modifiedRequest)
	if err != nil {
		fmt.Println("Error storing URL in the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shorturl":  modifiedRequest.ShortURL,
		"longurl":   modifiedRequest.LongURL,
		"createdat": modifiedRequest.CreatedAt,
		"error":     nil,
	})
}

func GetLongURL(c *gin.Context) {
	shortenedurl := c.Param("encodedurl")

	urlData := url.URL{
		ShortURL:  shortenedurl,
		CreatedAt: time.Now(),
	}

	// open db
	db, err := common.GetDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error while beginning transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	longURL, err := url.GetURLWithTransaction(tx, urlData)
	if err != nil {
		// Handle the error, you can choose how to respond to different error types
		if strings.Contains(err.Error(), "Shortened URL not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shortened URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Redirect or respond with the longURL
	c.Redirect(http.StatusFound, longURL)
	// Alternatively, you can respond with JSON:
	// c.JSON(http.StatusOK, gin.H{"longurl": longURL})

}

// =======NOTES ON BINDINGJSON========
// Use ShouldBindJSON if you want to handle binding errors yourself and continue with the request flow.
// Use BindJSON if you prefer a simpler way to handle binding errors by responding with a 400 Bad Request status directly.
