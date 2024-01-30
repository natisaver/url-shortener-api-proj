package routes

import (
	"fmt"
	"net/http"
	"urlshortener/controllers"
	"urlshortener/models/urlmodel"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// request flow is from:
// > server
// > handler
// > controller
// > repo/data management(DM)/data access(CRUD)
// which handles models

//usually we put controller

func ShortenURL(c *gin.Context) {
	// Parse the JSON request body into a URL object
	var request urlmodel.URL
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
	modifiedRequest := urlmodel.URL{
		LongURL:  sanitizedURL,
		ShortURL: shortenedURL,
	}

	// Call the controller to handle business logic
	// Create controller instance
	controllerInstance := controllers.NewURLController()
	err = controllerInstance.StoreURLController(modifiedRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"shorturl": modifiedRequest.ShortURL,
		"longurl":  modifiedRequest.LongURL,
		"error":    nil,
	})
}

func GetLongURL(c *gin.Context) {
	shortenedurl := c.Param("encodedurl")

	urlData := urlmodel.URL{
		ShortURL: shortenedurl,
	}

	// Call the controller to handle business logic
	// Create controller instance
	controllerInstance := controllers.NewURLController()
	longURL, err := controllerInstance.GetLongURLController(urlData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return Response
	// Redirect or respond with the longURL
	c.Redirect(http.StatusFound, longURL)
	// Alternatively, you can respond with JSON:
	// c.JSON(http.StatusOK, gin.H{"longurl": longURL})

}

// =======NOTES ON BINDINGJSON========
// Use ShouldBindJSON if you want to handle binding errors yourself and continue with the request flow.
// Use BindJSON if you prefer a simpler way to handle binding errors by responding with a 400 Bad Request status directly.
