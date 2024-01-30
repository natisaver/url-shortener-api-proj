package controllers

// controller is business logic
// one layer below handler
// call repo.crud, create tx and commit
// NEVER put commit in the defer function as that will always run at the end
// anytime there is an error, you will want to rollback immediately
// but there may be situations where you only rollback upon certain errors
// when creating the tx, no need for rollback

// FORMAT =============
// func sth() {
// 	tx = db.GetTx
// 	// defer tx.Rollback()
// 	defer stopPanic(tx.Rollback)

// 	err := updateDBRec(...)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// }
// =====================

import (
	"fmt"
	"urlshortener/common"
	"urlshortener/models/urlmodel"
	repo "urlshortener/repo/urlrepo"
	"urlshortener/utils"

	_ "github.com/go-sql-driver/mysql"
)

type URLControllerInterface interface {
	StoreURLController(urlmodel.URL) error
	GetLongURLController(urlmodel.URL) (string, error)
}

type urlController struct{}

func NewURLController() URLControllerInterface {
	return &urlController{}
}

func (u *urlController) StoreURLController(requestObj urlmodel.URL) error {

	// Open the database
	// db, err := common.GetDB()
	db := common.GetDB()

	// if err != nil {
	// 	fmt.Println("Error opening database connection:", err)
	// 	return err
	// }
	// defer db.Close()

	// Create transaction, i.e., a temporary form of our DB connection
	// tx, err := db.Begin()
	tx := db.Begin()
	// if err != nil {
	// 	fmt.Println("Error beginning database transaction:", err)
	// 	return err
	// }

	// defer in case of any exceptions
	defer func() {
		err := utils.StopPanic(tx.Rollback)
		if err != nil {
			fmt.Println("Error from StopPanic:", err)
		}
	}()

	// Create CRUD repository instance
	repoInstance := repo.NewCRUDRepository(tx)

	// ensure all queries are in one transaction, ensuring consistency of data
	// instead of passing in the db connection obj

	// call crud
	// err := repo.(tx,requestObj)
	err := repoInstance.StoreURL(requestObj)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error storing URL in the database:", err)
		return err
	}

	//if no exception
	tx.Commit()
	return nil
}

func (u *urlController) GetLongURLController(requestObj urlmodel.URL) (string, error) {
	// open db
	// db, err := common.GetDB()
	// with ORM
	db := common.GetDB()
	tx := db.Begin()

	// defer in case of any exceptions
	defer func() {
		err := utils.StopPanic(tx.Rollback)
		if err != nil {
			fmt.Println("Error from StopPanic:", err)
		}
	}()

	// Create CRUD repository instance
	repoInstance := repo.NewCRUDRepository(tx)

	// call crud
	longURL, err := repoInstance.GetURL(requestObj)
	// longURL, err := repo.GetURL(tx, requestObj)
	if err != nil {
		// rollback immediately
		tx.Rollback()
		fmt.Println("Shortened url not found:", err)
		// Handle the error, you can choose how to respond to different error types
		return "", err
	}

	//if no exception
	tx.Commit()
	return longURL, nil
}
