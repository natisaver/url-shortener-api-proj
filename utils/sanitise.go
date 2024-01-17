package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func SanitizeURL(rawURL string) (string, error) {
	// Check if the URL has a scheme, if not, assume it's "http"
	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	// Parse the URL
	// If the parsing is successful, a url.URL structure (parsedURL) is obtained, representing the various components of the URL.
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Check if the host is a valid domain
	if !isValidDomain(parsedURL.Host) {
		return "", fmt.Errorf("invalid URL: invalid domain")
	}

	// Check if the parsed URL has a scheme and host
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL: missing scheme or host")
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

func isValidDomain(host string) bool {
	// Perform additional checks for a valid domain (you can customize this)
	domainRegex := regexp.MustCompile(`^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(host)
	// return strings.Contains(host, ".")
}
