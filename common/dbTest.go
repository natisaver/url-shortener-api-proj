package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func InitDBTest() (*gorm.DB, error) {
	// Create a connection string for GORM
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database using GORM
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Store the GORM DB instance
	testDB = dbInstance

	// Migrate the database schema here if needed
	// db.AutoMigrate(&YourModel{})

	fmt.Println("Successfully connected to DB!")

	// Return the GORM DB instance
	return dbInstance, nil
}

func GetDBTest() *gorm.DB {
	// Return the stored GORM DB instance
	return testDB
}
