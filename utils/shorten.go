package utils

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func ShortenURL(sanitizedURL string) (string, error) {
	// Sanitize and validate the input URL
	parsedURL, err := url.Parse(sanitizedURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return "", fmt.Errorf("Invalid URL: %s", sanitizedURL)
	}

	// Base64 encode the URL to create a short representation
	shortURL := base64.URLEncoding.EncodeToString([]byte(originalURL))

	// Remove padding characters from base64 encoding
	shortURL = strings.TrimRight(shortURL, "=")

	// You can prepend your local host or a custom domain
	// but im storing the encoded form directly
	// shortURL = "http://localhost/" + shortURL

	return shortURL, nil
}