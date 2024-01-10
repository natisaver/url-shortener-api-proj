package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"urlshortener/utils"
	"urlshortener/models/url"
	"urlshortener/common"
	"strings"
)

func ShortenURL(c *gin.Context) {
	// Parse the JSON request body into a URL object
	var request url.URL
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// sanitise for security
	sanitizedURL, err := utils.SanitizeURL(request.LongURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// shorten it to an encoded form
	shortenedURL, err := utils.ShortenURL(sanitizedURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new object with the modified shortURL
    modifiedRequest := url.URL{
        LongURL:   sanitizedURL,
        ShortURL:  shortenedURL,
		CreatedAt: time.Now(),
    }

	// store into db
	// we first create a transaction i.e. a temporary form of our DB connection
	tx, err := common.GetDB().Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	// instead of passing in the actual DB connection
	// this is to ensure all queries are in one transaction, ensuring consistency of data
	err = url.StoreURLWithTransaction(tx, modifiedRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"shortenedURL": modifiedRequest.ShortURL, "longURL": modifiedRequest.LongURL, "createdAt": modifiedRequest.CreatedAt, "error": nil})
}

func GetLongURL(c *gin.Context) {
	shortenedurl := c.Param("encodedurl")
	urlData := url.URL{
        ShortURL:  shortenedurl,
		CreatedAt: time.Now(),
    }
	tx, err := common.GetDB().Begin()
	if err != nil {
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