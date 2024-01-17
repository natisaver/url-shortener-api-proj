package utils

import (
	"fmt"
	"time"
	"urlshortener/common"

	_ "github.com/go-sql-driver/mysql"
)

func cleanupExpiredRecords() {
	fmt.Println("Clean up service intialising...")
	// without ORM
	// db, err := common.GetDB()
	// with ORM
	db, err := common.GetDB().DB()
	if err != nil {
		fmt.Println("Error opening database connection for cleanup:", err)
		return
	}
	defer db.Close()

	// Calculate the date 30 days ago
	thresholdDate := time.Now().Add(-30 * 24 * time.Hour)

	// Prepare the DELETE statement
	_, err = db.Exec("DELETE FROM urls WHERE createdat < $1", thresholdDate)
	if err != nil {
		fmt.Println("Error executing delete statement:", err)
		return
	}

	fmt.Println("Clean up service success...")
}

func RunCleanupJob() {
	fmt.Println("Run clean up job called")
	// Run the cleanup job every 24 hours
	for range time.Tick(24 * time.Hour) {
		cleanupExpiredRecords()
	}
}
