package url

import (
	"database/sql"
	"fmt"
	"time"
)

type URL struct {
	ID        string    `json:"id"`
	ShortURL  string    `json:"shorturl"`
	LongURL   string    `json:"longurl"`
	CreatedAt time.Time `json:"created_at"`
}

// StoreURLWithTransaction stores URL information into the database within a transaction
func StoreURLWithTransaction(tx *sql.Tx, url URL) error {
	// Declare err at the beginning of the function
	var err error

	// Defer a function to either commit or rollback the transaction
	// () at the end is used to invoke the function immediately
	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback the transaction
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				fmt.Println("Error rolling back transaction:", rollbackErr)
			}
			// Re-panic after rollback as it's caught by recover()
			panic(p)
		} else if err != nil {
			// An error occurred, rollback the transaction
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				fmt.Println("Error rolling back transaction:", rollbackErr)
			}
		} else {
			// All good, commit the transaction
			err = tx.Commit()
			if err != nil {
				// Handle the commit error or propagate it
				fmt.Println("Error committing transaction:", err)
			}
		}
	}()

	// Check if shorturl already exists within the transaction
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM urls WHERE shorturl = $1", url.ShortURL).Scan(&count)
	if err != nil {
		return err
	}

	// If count is greater than 0, shorturl already exists
	// dont return an error as we want the user to receive the already shortened url
	if count > 0 {
		// return fmt.Errorf("ShortURL '%s' already exists in the database", url.ShortURL)
		return nil
	}

	// If shorturl does not exist, insert the new record within the transaction
	_, err = tx.Exec("INSERT INTO urls (shorturl, longurl) VALUES ($1, $2)", url.ShortURL, url.LongURL)
	if err != nil {
		return err
	}

	return nil
}

func GetURLWithTransaction(tx *sql.Tx, url URL) (string, error) {
	var err error

	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback the transaction
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				fmt.Println("Error rolling back transaction:", rollbackErr)
			}
			// Re-panic after rollback as it's caught by recover()
			panic(p)
		} else if err != nil {
			// An error occurred, rollback the transaction
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				fmt.Println("Error rolling back transaction:", rollbackErr)
			}
		} else {
			// All good, commit the transaction
			err = tx.Commit()
			if err != nil {
				// Handle the commit error or propagate it
				fmt.Println("Error committing transaction:", err)
			}
		}
	}()

	// Query the database to find the longurl based on the shortenedURL
	var longURL string
	err = tx.QueryRow("SELECT longurl FROM urls WHERE shorturl = $1", url.ShortURL).Scan(&longURL)

	if err == sql.ErrNoRows {
		// Handle case where no matching record was found
		return "", fmt.Errorf("shortened URL not found")
	} else if err != nil {
		// Handle other database query errors
		return "", fmt.Errorf("internal Server Error")
	}

	return longURL, nil
}
