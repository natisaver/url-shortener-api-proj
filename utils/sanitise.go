package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func SanitizeURL(rawURL string) (string, error) {
	// Check if the URL has a scheme, if not, assume it's "http"
	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	// Parse the URL
	// If the parsing is successful, a url.URL structure (parsedURL) is obtained, representing the various components of the URL.
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", err
	}

	// Reassemble the URL with cleaned components
	// The scheme, host, and path are included in the sanitizedURL in the format %s://%s%s. This ensures that the URL has a consistent structure.
	sanitizedURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path)

	// Add query parameters if present ?
	// https://www.example.com/search?q=golang&category=programming
	if parsedURL.RawQuery != "" {
		sanitizedURL += "?" + parsedURL.RawQuery
	}

	// Add fragment if present #
	// https://www.example.com/page#section-1
	if parsedURL.Fragment != "" {
		sanitizedURL += "#" + parsedURL.Fragment
	}

	return sanitizedURL, nil
}
