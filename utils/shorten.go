package utils

import (
	"encoding/base64"
	"strings"
)

func ShortenURL(sanitizedURL string) (string, error) {

	// Base64 encode the URL to create a short representation
	shortURL := base64.URLEncoding.EncodeToString([]byte(sanitizedURL))

	// Remove padding characters from base64 encoding
	shortURL = strings.TrimRight(shortURL, "=")

	// You can prepend your local host or a custom domain
	// but im storing the encoded form directly
	// shortURL = "http://localhost/" + shortURL

	return shortURL, nil
}
