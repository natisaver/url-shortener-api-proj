package urlrepo

import (
	"fmt"
	"urlshortener/models/urlmodel"

	"gorm.io/gorm"
)

// Create new URL record
func StoreURL(tx *gorm.DB, model urlmodel.URL) error {
	fmt.Println("CRUD CALLED IN")
	fmt.Println(model)
	// Check if shorturl already exists
	var count int64
	result := tx.Model(&urlmodel.URL{}).Where("shorturl = ?", model.ShortURL).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	// If count is greater than 0, shorturl already exists
	// don't return an error as we want the user to receive the already shortened url
	if count > 0 {
		return nil
	}

	// If shorturl does not exist, insert the new record
	result = tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Retrieve URL record
func GetURL(db *gorm.DB, model urlmodel.URL) (string, error) {
	fmt.Println("CRUD CALLED")
	var longURL string
	result := db.Model(&urlmodel.URL{}).Select("longurl").Where("shorturl = ?", model.ShortURL).Scan(&longURL)
	if result.Error == gorm.ErrRecordNotFound {
		// Handle case where no matching record was found
		return "", fmt.Errorf("shortened URL not found")
	} else if result.Error != nil {
		// Handle other database query errors
		return "", fmt.Errorf("internal server error")
	}

	return longURL, nil
}

// ========== OLD CODE changed because: =========================================================================
// 1) changed to orm, 2) tx.commit should be in the controller layer 3) implement external catch panic function
// ==============================================================================================================

// StoreURLWithTransaction stores URL information into the database within a transaction

// // Create new URL record
// func StoreURL(tx *sql.Tx, url url.URL) error {

// 	// Check if shorturl already exists within the transaction
// 	var count int
// 	err := tx.QueryRow("SELECT COUNT(*) FROM urls WHERE shorturl = $1", url.ShortURL).Scan(&count)
// 	if err != nil {
// 		return err
// 	}

// 	// If count is greater than 0, shorturl already exists
// 	// dont return an error as we want the user to receive the already shortened url
// 	if count > 0 {
// 		// return fmt.Errorf("ShortURL '%s' already exists in the database", url.ShortURL)
// 		return nil
// 	}

// 	// If shorturl does not exist, insert the new record within the transaction
// 	_, err = tx.Exec("INSERT INTO urls (shorturl, longurl, createdat) VALUES ($1, $2, $3)", url.ShortURL, url.LongURL, url.CreatedAt)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Retrieve URL record
// func GetURL(tx *sql.Tx, url url.URL) (string, error) {
// 	var err error

// 	// Query the database to find the longurl based on the shortenedURL
// 	var longURL string
// 	err = tx.QueryRow("SELECT longurl FROM urls WHERE shorturl = $1", url.ShortURL).Scan(&longURL)

// 	if err == sql.ErrNoRows {
// 		// Handle case where no matching record was found
// 		return "", fmt.Errorf("shortened URL not found")
// 	} else if err != nil {
// 		// Handle other database query errors
// 		return "", fmt.Errorf("internal Server Error")
// 	}

// 	return longURL, nil
// }
