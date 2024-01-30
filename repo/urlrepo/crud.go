package urlrepo

import (
	"errors"
	"fmt"
	"urlshortener/models/urlmodel"

	"gorm.io/gorm"
)

// interface to show the methods
type CRUDRepositoryInterface interface {
	StoreURL(model urlmodel.URL) error
	GetURL(model urlmodel.URL) (string, error)
}

// struct for the repository
type crudRepository struct {
	tx *gorm.DB
}

// constructor
// needed since struct is private
func NewCRUDRepository(db *gorm.DB) CRUDRepositoryInterface {
	return &crudRepository{tx: db}
}

// Implementing methods of CRUDRepositoryInterface

// Create new URL record
func (c *crudRepository) StoreURL(model urlmodel.URL) error {
	if model == (urlmodel.URL{}) {
		return errors.New("empty object")
	}

	// Check if shorturl already exists
	var count int64
	result := c.tx.Model(&urlmodel.URL{}).Where("shorturl = ?", model.ShortURL).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	// If count is greater than 0, shorturl already exists
	// don't return an error as we want the user to receive the already shortened url
	if count > 0 {
		return nil
	}

	// If shorturl does not exist, insert the new record
	result = c.tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Retrieve URL record
func (c *crudRepository) GetURL(model urlmodel.URL) (string, error) {
	var longURL string

	if model == (urlmodel.URL{}) {
		return "", errors.New("empty object")
	}

	result := c.tx.Model(&urlmodel.URL{}).Select("longurl").Where("shorturl = ?", model.ShortURL).Scan(&longURL)
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

// package urlrepo

// import (
// 	"fmt"
// 	"urlshortener/models/urlmodel"

// 	"gorm.io/gorm"
// )

// // Create new URL record
// func StoreURL(tx *gorm.DB, model urlmodel.URL) error {
// 	fmt.Println("CRUD CALLED IN")
// 	fmt.Println(model)
// 	// Check if shorturl already exists
// 	var count int64
// 	result := tx.Model(&urlmodel.URL{}).Where("shorturl = ?", model.ShortURL).Count(&count)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	// If count is greater than 0, shorturl already exists
// 	// don't return an error as we want the user to receive the already shortened url
// 	if count > 0 {
// 		return nil
// 	}

// 	// If shorturl does not exist, insert the new record
// 	result = tx.Create(&model)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }

// // Retrieve URL record
// func GetURL(db *gorm.DB, model urlmodel.URL) (string, error) {
// 	fmt.Println("CRUD CALLED")
// 	var longURL string
// 	result := db.Model(&urlmodel.URL{}).Select("longurl").Where("shorturl = ?", model.ShortURL).Scan(&longURL)
// 	if result.Error == gorm.ErrRecordNotFound {
// 		// Handle case where no matching record was found
// 		return "", fmt.Errorf("shortened URL not found")
// 	} else if result.Error != nil {
// 		// Handle other database query errors
// 		return "", fmt.Errorf("internal server error")
// 	}

// 	return longURL, nil
// }
