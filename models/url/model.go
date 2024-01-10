package url

import (
	"database/sql"
	"fmt"
)

type URL struct {
    ID			string	`json:"id"`
    ShortURL	string  `json:"shorturl"`
    LongURL		string	`json:"longurl"`
	CreatedAt time.Time `json:"created_at"`
}

// StoreURLWithTransaction stores URL information into the database within a transaction
func StoreURLWithTransaction(tx *sql.DB, url URL) error {
	// Defer a function to either commit or rollback the transaction
	// () at the end is used to invoke the function immediately
	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback the transaction
			tx.Rollback()
			panic(p) // Re-panic after rollback as its caught by recover()
		} else if err != nil {
			// An error occurred, rollback the transaction
			tx.Rollback()
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

	// If count is greater than 0, shorturl already exists, return an error
	if count > 0 {
		return fmt.Errorf("ShortURL '%s' already exists in the database", url.ShortURL)
	}

	// If shorturl does not exist, insert the new record within the transaction
	_, err = tx.Exec("INSERT INTO urls (shorturl, longurl) VALUES ($1, $2)", url.ShortURL, url.LongURL)
	if err != nil {
		return err
	}

	return nil
}

func GetURLWithTransaction(tx *sql.DB, url URL) error {
	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, rollback the transaction
			tx.Rollback()
			panic(p) // Re-panic after rollback as its caught by recover()
		} else if err != nil {
			// An error occurred, rollback the transaction
			tx.Rollback()
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

	err := db.QueryRow("SELECT longurl FROM urls WHERE shorturl = ?", url.ShortURL).Scan(&longURL)

	if err == sql.ErrNoRows {
		// Handle case where no matching record was found
		c.JSON(http.StatusNotFound, gin.H{"error": "Shortened URL not found"})
		return
	} else if err != nil {
		// Handle other database query errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Redirect or respond with the longURL
	c.Redirect(http.StatusFound, longURL)
	// c.JSON(http.StatusOK, gin.H{"longurl": longURL})
}

