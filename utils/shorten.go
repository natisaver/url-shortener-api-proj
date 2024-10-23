package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ShortenURL(sanitizedURL string) (string, error) {
	// Generate MD5 hash of the sanitizedURL
	hasher := md5.New()
	hasher.Write([]byte(sanitizedURL))
	hashBytes := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	shortURL := hex.EncodeToString(hashBytes)

	// truncate to 8 characters
	// 16^8 = 4.29 billion urls
	shortURL = shortURL[:8]

	return shortURL, nil
}
