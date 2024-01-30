package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// func InitDBTest() (*gorm.DB, error) {
// 	// Create a connection string for GORM
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, "testdb")

// 	// Open a connection to the database using GORM
// 	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Store the GORM DB instance
// 	testDB = dbInstance

// 	// Migrate the database schema here if needed
// 	// dbInstance.AutoMigrate(&YourModel{})

// 	fmt.Println("Successfully connected to test DB!")

// 	// Return the GORM DB instance
// 	return dbInstance, nil
// }

// func CloseDBTest(db *gorm.DB) {
// 	// Close the test database connection
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		fmt.Println("Error getting SQL DB instance:", err)
// 	}
// 	sqlDB.Close()
// }

// // This function is called by the testing framework before running any tests
// func TestMain(m *testing.M) {
// 	// Run the tests
// 	exitCode := m.Run()

// 	// Cleanup and close the test database connection after all tests have run
// 	CloseDBTest(testDB)

// 	// Exit with the status code from the tests
// 	os.Exit(exitCode)
// }

// func CloseDBTest(db *gorm.DB) {
// 	// Close the test database connection
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		fmt.Println("Error getting SQL DB instance:", err)
// 	}
// 	sqlDB.Close()
// }

// // This function is called by the testing framework before running any tests
// func TestMain(m *testing.M) {
// 	// Run the tests
// 	exitCode := m.Run()

// 	// Cleanup and close the test database connection after all tests have run
// 	CloseDBTest(testDB)

// 	// Exit with the status code from the tests
// 	os.Exit(exitCode)
// }

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
