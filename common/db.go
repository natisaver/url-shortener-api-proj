package common

import (
	"fmt"
	"urlshortener/config"

	// Default SQL Driver
	// the _ imports the package without having to directly reference it in code
	// here it ensures the init() function of the postgre driver is called
	// "database/sql"
	// _ "github.com/lib/pq"

	// ORM
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host       = config.Host
	port       = config.Port
	user       = config.User
	password   = config.Password
	dbname     = config.Dbname
	dbtestname = config.Dbtestname
)

// We have 2 choices here for DB connection
// ORM => e.g. gorm dont have to write full sql statement, use functions instead (industry standard)
// SQL => need to write exact full sql statements

// with ORM

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
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
	db = dbInstance

	// Migrate the database schema here if needed
	// db.AutoMigrate(&YourModel{})

	fmt.Println("Successfully connected to DB!")

	// Return the GORM DB instance
	return dbInstance, nil
}

// we set the name of this function as a variable
// so that this variable can replaced with a testDB instead if necessary for testing
// other option is to pass a test DB via context for testing
// same concept as using a global time variable, var TimeNow = time.Now()
var GetDB = func() *gorm.DB {
	// Return the stored GORM DB instance
	return db
}

// without ORM

// func InitDB() (*sql.DB, error) {
// 	// SSL is chosen as disabled
// 	// likely to run into errors as its not defaultly enabled on lib/pq
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// validate details
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// defer means this runs before InitDB() returns
// 	// close the db connection after testing connectivity
// 	// release resource pool as there is limited access to the DB
// 	defer db.Close()

// 	// open db connection to test connectivity
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// store it for getter
// 	// dbobj = db

// 	fmt.Println("Successfully connected to DB!")

// 	// return a db connection
// 	// that you can then use to connect or disconnect from
// 	return db, nil
// }

// // public function
// // utilise singleton principle
// // getter for the initialised db connection
// func GetDB() (*sql.DB, error) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	return sql.Open("postgres", psqlInfo)
// }
