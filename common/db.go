package common

import (
	"database/sql"
	"fmt"
	// the _ imports the package without having to directly reference it in code
	// here it ensures the init() function of the postgre driver is called
	_ "github.com/lib/pq"
  )
  
  const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "urlDB"
  )

// variable to use for getter
var dbobj *sql.DB

func InitDB() (*sql.DB, error) {
	// SSL is chosen as disabled 
	// likely to run into errors as its not defaultly enabled on lib/pq
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ 
	"password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	// validate details
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	
	// defer means this runs before InitDB() returns
	// close the db connection after testing connectivity
	// release resource pool as there is limited access to the DB
	defer db.Close()

	// open db connection to test connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	// store it for getter
	dbobj = db

	fmt.Println("Successfully connected to DB!")

	// return a db connection
	// that you can then use to connect or disconnect from
	return db, nil
}

// public function
// utilise singleton principle
// getter for the initialised db connection
func GetDB() *sql.DB {
	return dbobj
}



